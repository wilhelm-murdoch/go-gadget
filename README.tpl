# Gadget

![CI Status](https://github.com/wilhelm-murdoch/go-gadget/actions/workflows/ci.yml/badge.svg)
![Release Status](https://github.com/wilhelm-murdoch/go-gadget/actions/workflows/release.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/wilhelm-murdoch/go-gadget?status.svg)](https://pkg.go.dev/github.com/wilhelm-murdoch/go-gadget)
[![Go report](https://goreportcard.com/badge/github.com/wilhelm-murdoch/go-gadget)](https://goreportcard.com/report/github.com/wilhelm-murdoch/go-gadget)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

`gadget` is a tool that allows you to quickly inspect your Go source code. It's effectively a small layer of abstraction built on top of the Go AST package.

It _inspects_ your _go_ code, hence the name:
![Go-go Gadget!](gadget.png)
### Why?
I was working on [another project](https://github.com/wilhelm-murdoch/go-collection) of mine and thought to myself, "It would be nice if I didn't have to constantly update this readme file every time I made a change." So, I started digging around Go's AST package and came up with `gadget`.
### But, pkg.go.dev already does this ...
Yeah, I know. But, I didn't fully realise I writing a crappier version of pkg.go.dev until about 90% into the project.

* Maybe you don't want people to leave your repository to understand the basics of your package's API. 
* Maybe you want to present this data in a different, more personalised, format.
* Maybe you can use this to write a basic linter, or just learn more about Go AST.

It was fun to write and I use the tool almost daily. Perhaps you'll find it useful as well.
## Download & Install

Binary releases are regularly published for the most common operating systems and CPU architectures. These can be downloaded from the [releases page](https://github.com/wilhelm-murdoch/go-gadget/releases). Presentingly, `gadget` has been tested on, and compiled for, the following:
1. Windows on `386`, `arm`, `amd64`
2. MacOS ( `debian` ) on `amd64`, `arm64`
3. Linux on `386`, `arm`, `amd64`, `arm64`

Download the appropriate archive and unpack the binary into your machine's local `$PATH`.

## Usage
Once added to your machine's local `$PATH` you can invoke `gadget` like so:
```
$ gadget --help
NAME:
   gadget - inspect your code via a small layer of abstraction over Go's AST package

USAGE:
   gadget [global options] command [command options] [arguments...]

VERSION:
   v0.0.12

AUTHOR:
   Wilhelm Murdoch <wilhelm@devilmayco.de>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --source value, -s value    path to the target go source file, or directory containing go source files. (default: ".")
   --format value, -f value    the output format of the results as json, template or debug. (default: "json")
   --template value, -t value  if the template format is selected, this is the path to the template file to use. (default: "README.tpl")
   --help, -h                  show help (default: false)
   --version, -v               print only the version (default: false)

COPYRIGHT:
   (c) 2022 Wilhelm Codes ( https://wilhelm.codes )
```
### Example Mapping
One of the primary goals of ths project was to link example functions to their associated functions and methods. As I'm regenerating templates, I'd like to have the example body as well as the expected output to live alongside the associated function so I can easily reference it within the [Go template](./README.tpl).

### JSON
Invoking the command with no flags will result in `gadget` searching for `*.go` files by recursively walking through the present working directory. Results will be displayed as a JSON object following this structure:

- `packages`: a list of discovered packages.
  - `files`: any `*.go` file associated with the package.
    - `types`: discovered types, eg; structs, interfaces, etc...
      - `fields`: a list of fields associated with each type.
    - `functions`: functions and methods
      - `examples`: mapped example functions ( if any )
    - `values`: explicitly-declared values, eg; constants and variables

A full example of the JSON object can be found in [here](./sink/sink.json).

### Debug
When invoking `gadget` using the `--format debug` flag, you will get output representing all evaluated source code using `ast.Print(...)`. Use this to follow the structure of the AST.
```bash
$ gadget --source /path/to/my/project --format debug
... heaps of AST output ...
1298  .  .  15: *ast.FuncDecl {
1299  .  .  .  Doc: *ast.CommentGroup {
1300  .  .  .  .  List: []*ast.Comment (len = 1) {
1301  .  .  .  .  .  0: *ast.Comment {
1302  .  .  .  .  .  .  Slash: sink/sink.go:81:1
1303  .  .  .  .  .  .  Text: "// GetPrivate is an accessor method that returns a dark secret:"
1304  .  .  .  .  .  }
1305  .  .  .  .  }
1306  .  .  .  }
1307  .  .  .  Recv: *ast.FieldList {
1308  .  .  .  .  Opening: sink/sink.go:82:6
1309  .  .  .  .  List: []*ast.Field (len = 1) {
1310  .  .  .  .  .  0: *ast.Field {
1311  .  .  .  .  .  .  Names: []*ast.Ident (len = 1) {
1312  .  .  .  .  .  .  .  0: *ast.Ident {
1313  .  .  .  .  .  .  .  .  NamePos: sink/sink.go:82:7
1314  .  .  .  .  .  .  .  .  Name: "nst"
1315  .  .  .  .  .  .  .  .  Obj: *ast.Object {
1316  .  .  .  .  .  .  .  .  .  Kind: var
1317  .  .  .  .  .  .  .  .  .  Name: "nst"
1318  .  .  .  .  .  .  .  .  .  Decl: *(obj @ 1310)
1319  .  .  .  .  .  .  .  .  }
1320  .  .  .  .  .  .  .  }
1321  .  .  .  .  .  .  }
... heaps more AST output
```
### Template
Use Go's template engine, along with [sprig](https://masterminds.github.io/sprig/), to generate technical documents or readme files ( like this one! ).
```bash
$ gadget --format template --template README.tpl > README.md
```
Or, without the `--template ...` flag as it will use `README.tpl` as the default template if it exists in the starting directory:
```
$ gadget --format template > README.md
```
The best way to understand this is by viewing the following "kitchen sink" examples:
1. [`sink/README.tpl`](./sink/README.tpl) is a valid Go template.
2. [`sink/README.md`](./sink/README.md) was generated using the specified Go template.

## Build Locally

`gadget` makes use of Go's new generics support, so the minimum viable version of the language is `1.18.x`. Ensure your local development environment meets this single requirement before continuing. There are also several build flags used when compiling the binary. These populate the output of the `gadget --version` flag.
```bash
$ git clone git@github.com:wilhelm-murdoch/go-gadget.git
$ cd gadget
$ make build
$ ./gadget --version
Version: v99.99.99, Stage: local, Commit: 9932cf9fdc90c0d8223ef85a0fc1ddfa99c28f95, Date: 10-04-2022
```
### Testing
All major functionality of `gadget` has been covered by testing. You can run the tests, benchmarks and lints using the following set of `Makefile` targets:
- `make test`: run the local testing suite.
- `make lint`: run `staticcheck` on the local source files.
- `make bench`: run a series of benchmarks and output the results as `cpu.out`, `mem.out` and `trace.out`
- `make pprof-cpu`: run a local webserver on port `:8800` displaying CPU usage stats.
- `make pprof-mem`: run a local webserver on port `:8900` displaying memory usage stats. 
- `make trace`: view local tracing output from the benchmark run.
- `make coverage`: view testing code coverage for the local source files.

## API

While `gadget` is meant to be used as a CLI, there's no reason you can't make use of it as a library to integrate into your own tools. If you were wondering, yes, this readme file was generated by `gadget` itself. 

{{ range . }}{{ if eq .Name "gadget" }}

{{ range .Files.Items }}
### File `{{ .Path }}` {{ $path := .Name }}
{{ range .Types.Items }}
#### Type `{{ .Name }}`
* `{{ trim .Signature }}` [#]({{ $path }}#L{{ .LineStart }})
* `{{ $path }}:{{ .LineStart }}:{{ .LineEnd }}` [#]({{ $path }}#L{{ .LineStart }}-L{{ .LineEnd }})
{{ if .Fields.Length }}
Exported Fields:
{{ range .Fields.Items }}
{{ if .IsExported }}1. `{{ .Name }}`: {{ .Comment }} [#]({{ $path }}#L{{ .Line }}){{ end }}{{ end }}{{ end }}---{{ end }}
{{ range .Functions.Items }}
#### Function `{{ .Name }}`
* `{{ trim .Signature }}` [#]({{ $path }}#L{{ .LineStart }})
* `{{ $path }}:{{ .LineStart }}:{{ .LineEnd }}` [#]({{ $path }}#L{{ .LineStart }}-L{{ .LineEnd }})
{{ if .Doc }}
{{ .Doc | replace "\n" "" }}{{ end }}
{{ if .Examples.Length }}
**Examples:**
{{ range .Examples.Items }}
```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-gadget"
)

func main() {
{{ indent 4 .Body }}
}
```
```
{{ .Output }}
```{{ end }}
{{ end }}
---{{ end }}{{ end }}{{ end }}{{ end }}
## License
Copyright © {{ now | date "2006" }} [Wilhelm Murdoch](https://wilhelm.codes).

This project is [MIT](./LICENSE) licensed.