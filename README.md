# Gadget

![CI Status](https://github.com/wilhelm-murdoch/go-gadget/actions/workflows/ci.yml/badge.svg)
![Release Status](https://github.com/wilhelm-murdoch/go-gadget/actions/workflows/release.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/wilhelm-murdoch/go-gadget?status.svg)](https://pkg.go.dev/github.com/wilhelm-murdoch/go-gadget)
[![Go report](https://goreportcard.com/badge/github.com/wilhelm-murdoch/go-gadget)](https://goreportcard.com/report/github.com/wilhelm-murdoch/go-gadget)

Gadget is a simple CLI that allows you to inspect your Go source code using a relatively flat and simple representation. It's effectively a simpler layer of abstraction built on top of the Go AST package.

It _inspects_ your _go_ code, hence the name:
![Go-go Gadget!](gadget.png)

## Download & Install

We publish binary releases for the most common operating systems and CPU architectures. These can be downloaded from the [releases page](https://github.com/wilhelm-murdoch/go-gadget/releases). Presentingly, Gadget has been tested on, and compiled for, the following:
1. Windows on `386`, `arm`, `amd64`
2. MacOS ( `debian` ) on `amd64`, `arm64`
3. Linux on `386`, `arm`, `amd64`, `arm64`

Simply download the appropriate archive and unpack the binary into your machine's local `$PATH`.

## Build Locally

## Usage

### Debug
### Template
### JSON

## API

While gadget is meant to be used as a CLI, there's no reason you can't make use of it as a library to intgrate into your own tools.

# License
Copyright © 2022 [Wilhelm Murdoch](https://wilhelm.codes).

This project is [MIT](./LICENSE) licensed.
