# Collection

![Build Status](https://github.com/wilhelm-murdoch/go-collection/actions/workflows/go.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/wilhelm-murdoch/go-collection?status.svg)](https://pkg.go.dev/github.com/wilhelm-murdoch/go-collection)
[![Go report](https://goreportcard.com/badge/github.com/wilhelm-murdoch/go-collection)](https://goreportcard.com/report/github.com/wilhelm-murdoch/go-collection)

A generic collection for Go with a few convenient methods.
# Install
```
go get github.com/wilhelm-murdoch/go-collection
```
{{ range . }}{{ range .Files.Items }}{{ $path := .Path }}{{ range .Functions.Items }}{{ if and (ne .Name "main") (not .IsTest) (not .IsExample) (not .IsBenchmark) (.IsExported)}}### Func {{ .Name }}
* `{{ trim .Signature }}` [#]({{ $path }}#L{{ .LineStart }})
* `{{ $path }}:{{ .LineStart }}-{{ .LineEnd }}` [#]({{ $path }}#L{{ .LineStart }}-L{{ .LineEnd }})

{{ .Comment | replace "\n" " " }}
{{ end }}{{ end }}{{ end }}{{ end }}
# License
Copyright © {{ now | date "2006" }} [Wilhelm Murdoch](https://wilhelm.codes).

This project is [MIT](./LICENSE) licensed.
