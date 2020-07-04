/*
Copyright 2020 The veneer Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package veneer

import (
	"archive/tar"
	"io"

	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// A FileOperator performs operations on a file in a tarball.
type FileOperator func(*tar.Header, io.Reader) error

// Walk walks the filesystem of image layers.
func Walk(layers []v1.Layer, f FileOperator) error {
	for _, layer := range layers {
		u, err := layer.Uncompressed()
		if err != nil {
			return err
		}
		defer u.Close()
		tr := tar.NewReader(u)
		for {
			h, err := tr.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			if h.FileInfo().IsDir() {
				continue
			}
			if err := f(h, tr); err != nil {
				return err
			}
		}
	}
	return nil
}
