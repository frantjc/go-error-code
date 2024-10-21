# go-error-code [![CI](https://github.com/frantjc/go-error-code/actions/workflows/ci.yml/badge.svg?branch=main&event=push)](https://github.com/frantjc/go-error-code/actions) [![godoc](https://pkg.go.dev/badge/github.com/frantjc/go-error-code.svg)](https://pkg.go.dev/github.com/frantjc/go-error-code) [![goreportcard](https://goreportcard.com/badge/github.com/frantjc/go-error-code)](https://goreportcard.com/report/github.com/frantjc/go-error-code) ![license](https://shields.io/github/license/frantjc/go-error-code)

Opinionated [Go](https://go.dev) module for bubbling up and consuming exit codes alongside errors. Useful for packages that are used by CLIs to return helpful codes to the caller without having to parse error strings.

## install

```sh
go get github.com/frantjc/go-error-code
```
