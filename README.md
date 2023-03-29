# A collection of useful libraries

This repository contains a collection of useful libraries that I have been building over time. The libraries are used across my private Golang projects.
I hope you find joy in using them as much as I do.
Each package is meant to be standalone and does not have dependencies on other packages.

## Versioning

- All tags in the repository use the nomenclature " v<major>.<minor>.<patch> ".
- Using the 'v' as suffix is important for publishing new versions of the module.

## Publishing

This library is published as a public library. Command:
`$ GOPROXY=proxy.golang.org go list -m github.com/tukaianirban/sdk.go.common@v1.2.5`

### Notes

- The jsonreader package will be soon deprecated.
