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

package main

import (
	"fmt"
	"os"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/spf13/afero"

	"github.com/hasheddan/veneer"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	ref, err := name.ParseReference(os.Args[1])
	if err != nil {
		panic(err)
	}

	img, err := remote.Image(ref, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		panic(err)
	}

	// Create in memory filesystem.
	fs := afero.NewMemMapFs()
	err = veneer.ImageFs(img, fs)
	if err != nil {
		panic(err)
	}
	if err = afero.Walk(fs, ".", func(path string, _ os.FileInfo, err error) error {
		fmt.Println(path)
		return err
	}); err != nil {
		panic(err)
	}

	layers, err := img.Layers()
	if err != nil {
		panic(err)
	}

	fmt.Println("-----------------")

	// Only look at filesystem in top layer.
	fs = afero.NewMemMapFs()
	err = veneer.LayerFs(layers[len(layers)-1], fs)
	if err != nil {
		panic(err)
	}
	if err = afero.Walk(fs, ".", func(path string, _ os.FileInfo, err error) error {
		fmt.Println(path)
		return err
	}); err != nil {
		panic(err)
	}
}
