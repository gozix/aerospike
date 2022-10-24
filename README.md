# GoZix Aerospike

[documentation-img]: https://img.shields.io/badge/godoc-reference-blue.svg?color=24B898&style=for-the-badge&logo=go&logoColor=ffffff
[documentation-url]: https://pkg.go.dev/github.com/gozix/aerospike/v2
[license-img]: https://img.shields.io/github/license/gozix/aerospike.svg?style=for-the-badge
[license-url]: https://github.com/gozix/aerospike/blob/master/LICENSE
[release-img]: https://img.shields.io/github/tag/gozix/aerospike.svg?label=release&color=24B898&logo=github&style=for-the-badge
[release-url]: https://github.com/gozix/aerospike/releases/latest
[build-status-img]: https://img.shields.io/github/actions/workflow/status/gozix/aerospike/go.yml?logo=github&style=for-the-badge
[build-status-url]: https://github.com/gozix/aerospike/actions
[go-report-img]: https://img.shields.io/badge/go%20report-A%2B-green?style=for-the-badge
[go-report-url]: https://goreportcard.com/report/github.com/gozix/aerospike
[code-coverage-img]: https://img.shields.io/codecov/c/github/gozix/aerospike.svg?style=for-the-badge&logo=codecov
[code-coverage-url]: https://codecov.io/gh/gozix/aerospike

[![License][license-img]][license-url]
[![Documentation][documentation-img]][documentation-url]

[![Release][release-img]][release-url]
[![Build Status][build-status-img]][build-status-url]
[![Go Report Card][go-report-img]][go-report-url]
[![Code Coverage][code-coverage-img]][code-coverage-url]

The bundle provide aerospike integration to GoZix application.

## Installation

```shell
go get github.com/gozix/aerospike/v2
```

## Dependencies

* [viper](https://github.com/gozix/viper)

## Configuration example

```json
{
  "aerospike_cluster": {
    "nodes": [
      "127.0.0.1:3000"
    ],
    "idle_timeout": "5m"
  }
}
```

## Documentation

You can find documentation on [pkg.go.dev][documentation-url] and read source code if needed.

## Questions

If you have any questions, feel free to create an issue.
