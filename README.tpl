# Collection

![Build Status](https://github.com/wilhelm-murdoch/go-collection/actions/workflows/go.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/wilhelm-murdoch/go-collection?status.svg)](https://pkg.go.dev/github.com/wilhelm-murdoch/go-collection)
[![Go report](https://goreportcard.com/badge/github.com/wilhelm-murdoch/go-collection)](https://goreportcard.com/report/github.com/wilhelm-murdoch/go-collection)

A generic collection for Go with a few convenient methods.
# Install
```
go get github.com/wilhelm-murdoch/go-collection
```
# Usage
Import `go-collection` with the following:
```go
package main

import (
  "fmt"

  "github.com/wilhelm-murdoch/go-collection"
)

func main() {
	fruits := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot")
	fmt.Println("Fruits:", fruits.Length())

	fruits.Each(func(index int, item string) bool {
		fmt.Println("-", item)
		return false
	})

	// Output:
	// Fruits: 6
	// - apple
	// - orange
	// - strawberry
	// - cherry
	// - banana
	// - apricot
}
```
# Methods
{{ range .Functions }}{{ if and (not (hasPrefix "Example" .Name)) (not (hasPrefix "Test" .Name)) }}
  * [{{ .Name }}](#{{ .Name }}){{ end }}{{ end }}
{{ range .Functions }}{{ if and (not (hasPrefix "Example" .Name)) (not (hasPrefix "Test" .Name)) (.Comment)}}
## {{ .Name }}
{{ .Comment }} {{ if .Example }}```go
{{ .Example }}
```
{{ end }}{{ end }}{{ end }}
# License
Copyright © {{ now | date "2006" }} [Wilhelm Murdoch](https://wilhelm.codes).

This project is [MIT](./LICENSE) licensed.