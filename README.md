# go-polynym
**go-polynym** is the unofficial golang implementation for the [Polynym API](https://polynym.io/)

[![Build Status](https://travis-ci.com/mrz1836/go-polynym.svg?branch=master&v=2)](https://travis-ci.com/mrz1836/go-polynym)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-polynym?style=flat&v=2)](https://goreportcard.com/report/github.com/mrz1836/go-polynym)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/85aed3f384894abc958e9fa1e7f2f7ac)](https://www.codacy.com/app/mrz1818/go-polynym?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mrz1836/go-polynym&amp;utm_campaign=Badge_Grade)
[![Release](https://img.shields.io/github/release-pre/mrz1836/go-polynym.svg?style=flat&v=1)](https://github.com/mrz1836/go-polynym/releases)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme)
[![GoDoc](https://godoc.org/github.com/mrz1836/go-polynym?status.svg&style=flat)](https://godoc.org/github.com/mrz1836/go-polynym)

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

**go-polynym** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy) and [dep](https://github.com/golang/dep).
```bash
$ go get -u github.com/mrz1836/go-polynym
```

Updating dependencies in **go-polynym**:
```bash
$ cd ../go-polynym
$ dep ensure -update -v
```

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/mrz1836/go-polynym).

### Features
- Resolves [RelayX Handles](https://relayx.io),  [$handcash handles](https://handcash.io), [Paymails](https://bsvalias.org/), and BitcoinSV addresses
- Client is completely configurable
- Customize User Agent per request
- Using [heimdall http client](https://github.com/gojek/heimdall) with exponential backoff & more

## Examples & Tests
All unit tests and [examples](polynym_test.go) run via [Travis CI](https://travis-ci.org/mrz1836/go-polynym) and uses [Go version 1.13.x](https://golang.org/doc/go1.13). View the [deployment configuration file](.travis.yml).

Run all tests (including integration tests)
```bash
$ cd ../go-polynym
$ go test ./... -v
```

Run tests (excluding integration tests)
```bash
$ cd ../go-polynym
$ go test ./... -v -test.short
```

## Benchmarks
Run the Go [benchmarks](polynym_test.go):
```bash
$ cd ../go-polynym
$ go test -bench . -benchmem
```

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

## Usage
- View the [polynym examples](#examples--tests) above

Basic implementation:
```golang
package main

import (
	"log"

	"github.com/mrz1836/go-polynym"
)

func main() {

	// Start a new client and resolve
	client, _ := polynym.NewClient()
	resp, _ := client.GetAddress("mrz@moneybutton.com")

	log.Println("address:", resp.Address)
}
```

## Maintainers

[@MrZ](https://github.com/mrz1836)

## Contributing

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-polynym)

#### Credits

[@Dean](https://github.com/deanmlittle) & [UptimeSV](https://github.com/uptimesv) for their hard work on the [Polynym project](https://polynym.io/)

Looking for a Javascript version? Check out the [Polynym npm package](https://www.npmjs.com/package/polynym)

## License

![License](https://img.shields.io/github/license/mrz1836/go-polynym.svg?style=flat&v=2)