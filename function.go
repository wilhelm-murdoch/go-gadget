package gadget

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"log"
	"regexp"
	"strings"
)

// Function represents a golang function or method along with meaningful fields.
type Function struct {
	Name        string      `json:"name"`
	IsTest      bool        `json:"is_test"`
	IsBenchmark bool        `json:"is_benchmark"`
	IsExample   bool        `json:"is_example"`
	IsExported  bool        `json:"is_exported"`
	Examples    []*Function `json:"examples"`
	Comment     string      `json:"comment"`
	Body        string      `json:"body"`
	Output      string      `json:"output"`
	Signature   string      `json:"signature"`
	LineStart   int         `json:"line_start"`
	LineEnd     int         `json:"line_end"`
	LineCount   int         `json:"line_count"`
	astFile     *ast.File
	astFunc     *ast.FuncDecl
	tokenSet    *token.FileSet
}

// NewFunction returns a function instance and attempts to populate all
// associated fields with meaningful values.
func NewFunction(fn *ast.FuncDecl, ts *token.FileSet, f *ast.File) *Function {
	return (&Function{
		Name:     fn.Name.Name,
		Comment:  fn.Doc.Text(),
		astFunc:  fn,
		tokenSet: ts,
		astFile:  f,
	}).Parse()
}

// Parse is responsible for browsing through f.astFunc, f.tokenSet and f.astFile
// to populate the current function's fields. ( Chainable )
func (f *Function) Parse() *Function {
	if strings.HasPrefix(f.Name, "Example") {
		f.IsExample = true
	} else if strings.HasPrefix(f.Name, "Benchmark") {
		f.IsBenchmark = true
	} else if strings.HasPrefix(f.Name, "Test") {
		f.IsTest = true
	}

	f.IsExported = f.astFunc.Name.IsExported()

	f.parseLines()
	f.parseBody()
	f.parseSignature()

	if f.IsExample {
		f.parseOutput()
	}

	return f
}

// parseOutput attempts to fetch the expected output block for an example
// function and pins it to the current Function for future reference. It assumes
// all comments following the position of string "// Output:" belong to the
// output block.
func (f *Function) parseOutput() {
	var outputPos token.Pos
	if f.astFile.Comments != nil {
		for _, cg := range f.astFile.Comments {
			for _, c := range cg.List {
				if c.Pos() >= f.astFunc.Pos() && c.End() <= f.astFunc.End() {
					if strings.HasSuffix(c.Text, "// Output:") {
						outputPos = c.Pos()
					}

					if c.Pos() >= outputPos {
						f.Output += fmt.Sprintf("%s\n", c.Text)
					}
				}
			}
		}
	}

	f.Output = strings.TrimSpace(f.Output)
}

// parseLines determines the current function body's line positions within the
// currently evaluated file.
func (f *Function) parseLines() {
	f.LineStart = f.tokenSet.File(f.astFunc.Pos()).Line(f.astFunc.Pos())
	f.LineEnd = f.tokenSet.Position(f.astFunc.End()).Line
	f.LineCount = f.LineEnd - f.LineStart
}

// parseBody attempts to make a few adjustments to the *ast.BlockStmt which
// represents the current function's body. We remove the opening and closing
// braces as well as the first occurrent \t sequence on each line.
func (f *Function) parseBody() {
	var b bytes.Buffer
	if err := format.Node(&b, f.tokenSet, f.astFunc.Body); err != nil {
		log.Println(err)
	}

	source := fmt.Sprintf("%s", b.Bytes())

	var pattern *regexp.Regexp

	// Remove first leading tab character:
	pattern = regexp.MustCompile(`(?m)^\t{1}`)
	source = pattern.ReplaceAllString(source, "")

	source = source[:len(source)-1] // Remove trailing } brace
	source = source[1:]             // Remove leading { brace

	f.Body = strings.TrimSpace(source) // Trim all leading and trailing space
}

// parseSignature attempts to determine the current function's type and assigns
// it to the Signature field of struct Function.
func (f *Function) parseSignature() {
	var b bytes.Buffer
	if err := format.Node(&b, f.tokenSet, f.astFunc.Type); err != nil {
		log.Println(err)
	}
	f.Signature += fmt.Sprintf("%s", b.Bytes())
}
