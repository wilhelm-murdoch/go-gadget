# Gadget

![CI Status](https://github.com/wilhelm-murdoch/go-gadget/actions/workflows/go.yml/badge.svg)
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
Yeah, I know. But, I didn't fully realise I was writing a crappier version of pkg.go.dev until about 90% into the project.

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
One of the primary goals of this project was to link example functions to their associated functions and methods. As I'm regenerating templates, I'd like to have the example body as well as the expected output to live alongside the associated function so I can easily reference it within the [Go template](./README.tpl).

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




### File `field.go` 

#### Type `Field`
* `type Field struct` [#](field.go#L11)
* `field.go:11:21` [#](field.go#L11-L21)

Exported Fields:

1. `Name`: The name of the field. [#](field.go#L12)
1. `IsExported`: Determines whether the field is exported. [#](field.go#L13)
1. `IsEmbedded`: Determines whether the field is an embedded type. [#](field.go#L14)
1. `Line`: The line number this field appears on in the associated source file. [#](field.go#L15)
1. `Signature`: The full definition of the field including name, arguments and return values. [#](field.go#L16)
1. `Comment`: Any inline comments associated with the field. [#](field.go#L17)
1. `Doc`: The comment block directly above this field's definition. [#](field.go#L18)

---

#### Function `NewField`
* `func NewField(f *ast.Field, parent *File) *Field` [#](field.go#L25)
* `field.go:25:33` [#](field.go#L25-L33)

NewField returns a field instance and attempts to populate all associatedfields with meaningful values.

---
#### Function `Parse`
* `func (f *Field) Parse() *Field` [#](field.go#L37)
* `field.go:37:56` [#](field.go#L37-L56)

Parse is responsible for browsing through f.astField, f.parent to populatethe current fields's fields. ( Chainable )

---
#### Function `parseSignature`
* `func (f *Field) parseSignature()` [#](field.go#L62)
* `field.go:62:70` [#](field.go#L62-L70)

parseSignature determines the position of the current field within theassociated source file and extracts the relevant line of code. We only wantthe content before any inline comments. This will also replace consecutivespaces with a single space.

---
#### Function `String`
* `func (f *Field) String() string` [#](field.go#L73)
* `field.go:73:75` [#](field.go#L73-L75)

String implements the Stringer struct and returns the current package's name.

---
### File `file.go` 

#### Type `File`
* `type File struct` [#](file.go#L14)
* `file.go:14:29` [#](file.go#L14-L29)

Exported Fields:

1. `Name`: The basename of the file. [#](file.go#L15)
1. `Path`: The full path to the file as specified by the caller. [#](file.go#L16)
1. `Package`: The name of the golang package associated with this file. [#](file.go#L17)
1. `IsMain`: Determines whether this file is part of package main. [#](file.go#L18)
1. `IsTest`: Determines whether this file is for golang tests. [#](file.go#L19)
1. `HasTests`: Determines whether this file contains golang tests. [#](file.go#L20)
1. `HasBenchmarks`: Determines whether this file contains benchmark tests. [#](file.go#L21)
1. `HasExamples`: Determines whether this file contains example tests. [#](file.go#L22)
1. `Imports`: A list of strings containing all the current file's package imports. [#](file.go#L23)
1. `Values`: A collection of declared golang values. [#](file.go#L24)
1. `Functions`: A collection of declared golang functions. [#](file.go#L25)
1. `Types`: A collection of declared golang types. [#](file.go#L26)

---

#### Function `NewFile`
* `func NewFile(path string) (*File, error)` [#](file.go#L34)
* `file.go:34:50` [#](file.go#L34-L50)

NewFile returns a file instance representing a file within a golang package.This function creates a new token set and parser instance representing thenew file's abstract syntax tree ( AST ).

**Examples:**

```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-gadget"
)

func main() {
    if file, err := gadget.NewFile("sink/sink.go"); err == nil {
    	file.Functions.Each(func(i int, function *gadget.Function) bool {
    		fmt.Printf("%s defined between lines %d and %d\n", function.Name, function.LineStart, function.LineEnd)
    		return false
    	})
    }
}
```
```
// Output:
// PrintVars defined between lines 30 and 34
// AssignCollection defined between lines 37 and 43
// PrintConst defined between lines 46 and 50
// NewNormalStructTest defined between lines 72 and 79
// GetPrivate defined between lines 82 and 84
// GetOccupation defined between lines 87 and 89
// GetFullName defined between lines 92 and 94
// notExported defined between lines 98 and 100
// NewGenericStructTest defined between lines 111 and 113
// GetPrivate defined between lines 116 and 118
// GetFullName defined between lines 121 and 123
// IsBlank defined between lines 126 and 126
```
```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-gadget"
)

func main() {
    if file, err := gadget.NewFile("sink/sink.go"); err == nil {
    	file.Types.Each(func(i int, t *gadget.Type) bool {
    		if t.Fields.Length() > 0 {
    			fmt.Printf("%s is a %s with %d fields:\n", t.Name, t.Kind, t.Fields.Length())
    			t.Fields.Each(func(i int, f *gadget.Field) bool {
    				fmt.Printf("- %s on line %d\n", f.Name, f.Line)
    				return false
    			})
    		}
    		return false
    	})
    }
}
```
```
// Output:
// InterfaceTest is a interface with 1 fields:
// - ImplementMe on line 54
// EmbeddedStructTest is a struct with 1 fields:
// - Occupation on line 59
// NormalStructTest is a struct with 5 fields:
// - First on line 64
// - Last on line 65
// - Age on line 66
// - private on line 67
// - EmbeddedStructTest on line 68
// GenericStructTest is a struct with 4 fields:
// - First on line 104
// - Last on line 105
// - Age on line 106
// - private on line 107
```
```go
package main

import (
  "fmt"
  "strings"

  "github.com/wilhelm-murdoch/go-gadget"
)

func main() {
    var buffer strings.Builder
    if file, err := gadget.NewFile("sink/sink.go"); err == nil {
    	encoder := json.NewEncoder(&buffer)
    	if err := encoder.Encode(file.Values.Items()); err != nil {
    		fmt.Println(err)
    	}
    }
    
    fmt.Println(buffer.String())
}
```
```
// Output:
// [{"kind":"const","name":"ONE","line":9,"body":"ONE   = 1 // represents the number 1"},{"kind":"const","name":"TWO","line":10,"body":"TWO   = 2 // represents the number 2"},{"kind":"const","name":"THREE","line":11,"body":"THREE = 3 // represents the number 3"},{"kind":"var","name":"one","line":16,"body":"one   = \"one\"   // represents the english spelling of 1"},{"kind":"var","name":"two","line":17,"body":"two   = \"two\"   // represents the english spelling of 2"},{"kind":"var","name":"three","line":18,"body":"three = \"three\" // represents the english spelling of 3"},{"kind":"var","name":"collection","line":27,"body":"var collection map[string]map[string]string // this should be picked up as an inline comment."}]
```

---
#### Function `Parse`
* `func (f *File) Parse() *File` [#](file.go#L56)
* `file.go:56:68` [#](file.go#L56-L68)

Parse is responsible for walking through the current file's abstract syntaxtree in order to populate it's fields. This includes imports, definedfunctions and methods, structs and interfaces and other declared values.( Chainable )

---
#### Function `parsePackage`
* `func (f *File) parsePackage()` [#](file.go#L71)
* `file.go:71:76` [#](file.go#L71-L76)

parsePackage updates the current file with package-related data.

---
#### Function `parseImports`
* `func (f *File) parseImports()` [#](file.go#L81)
* `file.go:81:85` [#](file.go#L81-L85)

parseImports is responsible for creating a list of package imports that havebeen defined within the current file and assinging them to the appropriateImports field.

---
#### Function `parseFunctions`
* `func (f *File) parseFunctions()` [#](file.go#L90)
* `file.go:90:106` [#](file.go#L90-L106)

parseFunctions is responsible for creating abstract representations offunctions and methods defined within the current file. All functions areadded to the Functions collection.

---
#### Function `parseTypes`
* `func (f *File) parseTypes()` [#](file.go#L111)
* `file.go:111:119` [#](file.go#L111-L119)

parseTypes is responsible for creating abstract representations of declaredgolang types defined within the current file. All findings are added to theTypes collection.

---
#### Function `parseValues`
* `func (f *File) parseValues()` [#](file.go#L124)
* `file.go:124:134` [#](file.go#L124-L134)

parseValues is responsible for creating abstract representations of variousgeneral values such as const and var blocks. All values are added to theValues collection.

---
#### Function `walk`
* `func (f *File) walk(fn func(ast.Node) bool)` [#](file.go#L138)
* `file.go:138:140` [#](file.go#L138-L140)

walk implements the walk interface which is used to step through syntaxtrees via a caller-supplied callback.

---
#### Function `String`
* `func (f *File) String() string` [#](file.go#L143)
* `file.go:143:145` [#](file.go#L143-L145)

String implements the Stringer struct and returns the current package's name.

---
#### Function `GetAstAttributes`
* `func (f *File) GetAstAttributes() (*ast.File, *token.FileSet)` [#](file.go#L149)
* `file.go:149:151` [#](file.go#L149-L151)

GetAstAttributes returns the values associated with the astFile and tokenSetprivate fields. This is typically used for debug mode.

---
### File `function.go` 

#### Type `Function`
* `type Function struct` [#](function.go#L16)
* `function.go:16:35` [#](function.go#L16-L35)

Exported Fields:

1. `Name`: The name of the function. [#](function.go#L18)
1. `IsTest`: Determines whether this is a test. [#](function.go#L19)
1. `IsBenchmark`: Determines whether this is a benchmark. [#](function.go#L20)
1. `IsExample`: Determines whether this is an example. [#](function.go#L21)
1. `IsExported`: Determines whether this function is exported. [#](function.go#L22)
1. `IsMethod`: Determines whether this a method. This will be true if this function has a receiver. [#](function.go#L23)
1. `Receiver`: If this method has a receiver, this field will refer to the name of the associated struct. [#](function.go#L24)
1. `Doc`: The comment block directly above this funciton's definition. [#](function.go#L25)
1. `Output`: If IsExample is true, this field should contain the comment block defining expected output. [#](function.go#L26)
1. `Body`: The body of this function; everything contained within the opening and closing braces. [#](function.go#L27)
1. `Signature`: The full definition of the function including receiver, name, arguments and return values. [#](function.go#L28)
1. `LineStart`: The line number in the associated source file where this function is initially defined. [#](function.go#L29)
1. `LineEnd`: The line number in the associated source file where the definition block ends. [#](function.go#L30)
1. `LineCount`: The total number of lines, including body, the interface occupies. [#](function.go#L31)
1. `Examples`: A list of example functions associated with the current function. [#](function.go#L32)

---

#### Function `NewFunction`
* `func NewFunction(fn *ast.FuncDecl, parent *File) *Function` [#](function.go#L39)
* `function.go:39:47` [#](function.go#L39-L47)

NewFunction returns a function instance and attempts to populate allassociated fields with meaningful values.

---
#### Function `Parse`
* `func (f *Function) Parse() *Function` [#](function.go#L51)
* `function.go:51:72` [#](function.go#L51-L72)

Parse is responsible for browsing through f.astFunc, f.tokenSet and f.astFileto populate the current function's fields. ( Chainable )

---
#### Function `parseReceiver`
* `func (f *Function) parseReceiver()` [#](function.go#L76)
* `function.go:76:95` [#](function.go#L76-L95)

parseReceiver attemps to assign the receiver of a method, if one even exists,and assigns it to the `Function.Receiver` field.

---
#### Function `parseOutput`
* `func (f *Function) parseOutput()` [#](function.go#L101)
* `function.go:101:121` [#](function.go#L101-L121)

parseOutput attempts to fetch the expected output block for an examplefunction and pins it to the current Function for future reference. It assumesall comments immediately following the position of string "// Output:"belong to the output block.

---
#### Function `parseLines`
* `func (f *Function) parseLines()` [#](function.go#L125)
* `function.go:125:129` [#](function.go#L125-L129)

parseLines determines the current function body's line positions within thecurrently evaluated file.

---
#### Function `parseBody`
* `func (f *Function) parseBody()` [#](function.go#L136)
* `function.go:136:145` [#](function.go#L136-L145)

parseBody attempts to make a few adjustments to the *ast.BlockStmt whichrepresents the current function's body. We remove the opening and closingbraces as well as the first occurrent `\t` sequence on each line. Some peoplewill ask, "wHy dOn't yOu uSe tHe aSt pAcKaGe fOr tHiS" to which I answer,"Because, I'm lazy. We have the file, we know which lines contain the body."

---
#### Function `parseSignature`
* `func (f *Function) parseSignature()` [#](function.go#L149)
* `function.go:149:152` [#](function.go#L149-L152)

parseSignature attempts to determine the current function's type and assignsit to the Signature field of struct Function.

---
#### Function `String`
* `func (f *Function) String() string` [#](function.go#L155)
* `function.go:155:157` [#](function.go#L155-L157)

String implements the Stringer struct and returns the current package's name.

---
### File `package.go` 

#### Type `Package`
* `type Package struct` [#](package.go#L9)
* `package.go:9:12` [#](package.go#L9-L12)

Exported Fields:

1. `Name`: The name of the current package. [#](package.go#L10)
1. `Files`: A collection of golang files associated with this package. [#](package.go#L11)---

#### Function `NewPackage`
* `func NewPackage(name string) *Package` [#](package.go#L16)
* `package.go:16:21` [#](package.go#L16-L21)

NewPackage returns a Package instance with an initialised collection used forassigning and iterating through files.

---
#### Function `String`
* `func (p *Package) String() string` [#](package.go#L24)
* `package.go:24:26` [#](package.go#L24-L26)

String implements the Stringer struct and returns the current package's name.

---
### File `type.go` 

#### Type `Type`
* `type Type struct` [#](type.go#L20)
* `type.go:20:33` [#](type.go#L20-L33)

Exported Fields:

1. `Name`: The name of the struct. [#](type.go#L21)
1. `Kind`: Determines the kind of type, eg; interface or struct. [#](type.go#L22)
1. `LineStart`: The line number in the associated source file where this struct is initially defined. [#](type.go#L23)
1. `LineEnd`: The line number in the associated source file where the definition block ends. [#](type.go#L24)
1. `LineCount`: The total number of lines, including body, the struct occupies. [#](type.go#L25)
1. `Comment`: Any inline comments associated with the struct. [#](type.go#L26)
1. `Doc`: The comment block directly above this struct's definition. [#](type.go#L27)
1. `Signature`: The full definition of the struct itself. [#](type.go#L28)
1. `Body`: The full body of the struct sourced directly from the associated file; comments included. [#](type.go#L29)
1. `Fields`: A collection of fields and their associated metadata. [#](type.go#L30)

---

#### Function `NewType`
* `func NewType(ts *ast.TypeSpec, parent *File) *Type` [#](type.go#L37)
* `type.go:37:45` [#](type.go#L37-L45)

NewType returns an struct instance and attempts to populate all associatedfields with meaningful values.

---
#### Function `Parse`
* `func (t *Type) Parse() *Type` [#](type.go#L49)
* `type.go:49:56` [#](type.go#L49-L56)

Parse is responsible for browsing through f.astSpec, f.astType, f.parent topopulate the current struct's fields. ( Chainable )

---
#### Function `parseLines`
* `func (t *Type) parseLines()` [#](type.go#L60)
* `type.go:60:64` [#](type.go#L60-L64)

parseLines determines the current struct's opening and closing linepositions.

---
#### Function `parseBody`
* `func (t *Type) parseBody()` [#](type.go#L69)
* `type.go:69:71` [#](type.go#L69-L71)

parseBody attempts to make a few adjustments to the *ast.BlockStmt whichrepresents the current struct's body. We remove the opening and closingbraces as well as the first occurrent `\t` sequence on each line.

---
#### Function `parseSignature`
* `func (t *Type) parseSignature()` [#](type.go#L75)
* `type.go:75:78` [#](type.go#L75-L78)

parseSignature attempts to determine the current structs's type and assignsit to the Signature field of struct Function.

---
#### Function `parseFields`
* `func (t *Type) parseFields()` [#](type.go#L82)
* `type.go:82:103` [#](type.go#L82-L103)

parseFields iterates through the struct's list of defined methods topopulate the Fields collection.

---
#### Function `String`
* `func (t *Type) String() string` [#](type.go#L106)
* `type.go:106:108` [#](type.go#L106-L108)

String implements the Stringer struct and returns the current package's name.

---
### File `util.go` 

#### Type `Walker`
* `type Walker func(ast.Node) boo` [#](util.go#L80)
* `util.go:80:80` [#](util.go#L80-L80)
---

#### Function `GetLinesFromFile`
* `func GetLinesFromFile(path string, from, to int) []byte` [#](util.go#L18)
* `util.go:18:46` [#](util.go#L18-L46)

GetLinesFromFile creates a byte reader for the file at the target path andreturns a slice of bytes representing the file content. This slice isrestricted the lines specified by the `from` and `to` arguments inclusively.This will return an empty byte if an empty file, or any error, is encountered.

---
#### Function `WalkGoFiles`
* `func WalkGoFiles(path string) (files []string)` [#](util.go#L51)
* `util.go:51:61` [#](util.go#L51-L61)

WalkGoFiles recursively moves through the directory tree specified by `path`providing a slice of files matching the `*.go` extention. Explicitlyspecifying a file will return that file.

---
#### Function `AdjustSource`
* `func AdjustSource(source string, adjustBraces bool) string` [#](util.go#L66)
* `util.go:66:77` [#](util.go#L66-L77)

AdjustSource is a convenience function that strips the opening and closingbraces of a function's ( or other things ) body and removes the first `\t`character on each remaining line.

---
#### Function `Visit`
* `func (w Walker) Visit(node ast.Node) ast.Visitor` [#](util.go#L83)
* `util.go:83:88` [#](util.go#L83-L88)

Visit steps through each node within the specified syntax tree.

---
### File `value.go` 

#### Type `Value`
* `type Value struct` [#](value.go#L9)
* `value.go:9:16` [#](value.go#L9-L16)

Exported Fields:

1. `Kind`: Describes the current value's type, eg; CONST or VAR. [#](value.go#L10)
1. `Name`: The name of the value. [#](value.go#L11)
1. `Line`: The line number within the associated source file in which this value was originally defined. [#](value.go#L12)
1. `Body`: The full content of the associated statement. [#](value.go#L13)

---

#### Function `NewValue`
* `func NewValue(id *ast.Ident, parent *File) *Value` [#](value.go#L19)
* `value.go:19:26` [#](value.go#L19-L26)

NewValue returns a Value instance.

---
#### Function `Parse`
* `func (v *Value) Parse() *Value` [#](value.go#L30)
* `value.go:30:35` [#](value.go#L30-L35)

Parse is responsible for browsing through f.astIdent and f.tokenSet topopulate the current value's fields. ( Chainable )

---
#### Function `String`
* `func (v *Value) String() string` [#](value.go#L38)
* `value.go:38:40` [#](value.go#L38-L40)

String implements the Stringer struct and returns the current package's name.

---
## License
Copyright © 2022 [Wilhelm Murdoch](https://wilhelm.codes).

This project is [MIT](./LICENSE) licensed.
