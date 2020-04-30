# go-polynym
> The unofficial Go implementation for the [Polynym API](https://polynym.io/)

[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-polynym)](https://golang.org/)
[![Build Status](https://travis-ci.com/mrz1836/go-polynym.svg?branch=master&v=2)](https://travis-ci.com/mrz1836/go-polynym)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-polynym?style=flat&v=2)](https://goreportcard.com/report/github.com/mrz1836/go-polynym)
[![Release](https://img.shields.io/github/release-pre/mrz1836/go-polynym.svg?style=flat&v=1)](https://github.com/mrz1836/go-polynym/releases)
[![GoDoc](https://godoc.org/github.com/mrz1836/go-polynym?status.svg&style=flat)](https://pkg.go.dev/github.com/mrz1836/go-polynym)

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Installation

**go-polynym** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/mrz1836/go-polynym
```

## Documentation
You can view the generated [documentation here](https://pkg.go.dev/github.com/mrz1836/go-polynym).

### Features
- Resolves [RelayX Handles](https://relayx.io),  [$handcash handles](https://handcash.io), [Paymails](https://bsvalias.org/), and BitcoinSV addresses
- [Client](client.go) is completely configurable
- Using [heimdall http client](https://github.com/gojek/heimdall) with exponential backoff & more

<details>
<summary><strong><code>Library Deployment</code></strong></summary>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
all                            Runs test, install, clean, docs
bench                          Run all benchmarks in the Go application
clean                          Remove previous builds and any test cache data
clean-mods                     Remove all the Go mod cache
coverage                       Shows the test coverage
godocs                         Sync the latest tag with GoDocs
help                           Show all make commands available
lint                           Run the Go lint application
release                        Full production release (creates release in Github)
release-test                   Full production test release (everything except deploy)
release-snap                   Test the full release (build binaries)
tag                            Generate a new tag and push (IE: make tag version=0.0.0)
tag-remove                     Remove a tag if found (IE: make tag-remove version=0.0.0)
tag-update                     Update an existing tag to current commit (IE: make tag-update version=0.0.0)
test                           Runs vet, lint and ALL tests
test-short                     Runs vet, lint and tests (excludes integration tests)
update                         Update all project dependencies
update-releaser                Update the goreleaser application
vet                            Run the Go vet application
```
</details>

## Examples & Tests
All unit tests and [examples](polynym_test.go) run via [Travis CI](https://travis-ci.org/mrz1836/go-polynym) and uses [Go version 1.14.x](https://golang.org/doc/go1.14). View the [deployment configuration file](.travis.yml).

Run all tests (including integration tests)
```shell script
make test
```

Run tests (excluding integration tests)
```shell script
make test-short
```

## Benchmarks
Run the Go [benchmarks](polynym_test.go):
```shell script
make bench
```

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

## Usage
- View the [polynym examples](#examples--tests) above

Basic implementation:
```go
package main

import (
	"log"

	"github.com/mrz1836/go-polynym"
)

func main() {

	// Start a new client and resolve
	client, _ := polynym.NewClient(nil)
	resp, _ := client.GetAddress("mrz@moneybutton.com")

	log.Println("address:", resp.Address)
}
```

## Maintainers

| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:---:|
| [MrZ](https://github.com/mrz1836) |

## Contributing

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-polynym)

#### Credits

[@Dean](https://github.com/deanmlittle) & [BitPing](https://github.com/uptimesv) for their hard work on the [Polynym project](https://polynym.io/)

Looking for a Javascript version? Check out the [Polynym npm package](https://www.npmjs.com/package/polynym).

## License

![License](https://img.shields.io/github/license/mrz1836/go-polynym.svg?style=flat&v=2)