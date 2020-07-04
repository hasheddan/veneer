# veneer

A library for representing [OCI](https://opencontainers.org/) image layers in an
abstract filesystem, without requiring the presence of a container runtime.

## Future

Currently `veneer` takes a naive approach of achieving compatibility with
[`afero`](https://github.com/spf13/afero) by copying files from layer tarballs
to whatever filesystem backend is passed. In the future the tarball itself could
be used as the filesystem backend to avoid the overhead of copying all files.

## Dependencies

`veneer` heavily relies on and is influenced by
[`afero`](https://github.com/spf13/afero) and
[`go-containerregistry`](https://github.com/google/go-containerregistry).

## License

`veneer` is under the Apache 2.0 license.
