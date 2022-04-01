package gadget

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"log"
	"strings"
)

// Function represents a golang function or method along with meaningful fields.
type Function struct {
	Name        string      `json:"name"`         // The name of the function.
	IsTest      bool        `json:"is_test"`      // Determines whether this is a test.
	IsBenchmark bool        `json:"is_benchmark"` // Determines whether this is a benchmark.
	IsExample   bool        `json:"is_example"`   // Determines whether this is an example.
	IsExported  bool        `json:"is_exported"`  // Determines whether this function is exported.
	Examples    []*Function `json:"examples"`     // Any example functions of this function will be added to this list.
	Comment     string      `json:"comment"`      // Any inline comments associated with the function.
	Doc         string      `json:"doc"`          // The comment block directly above this funciton's definition.
	Output      string      `json:"output"`       // If IsExample is true, this field will contain the comment block defining expected output.
	Body        string      `json:"body"`         // The body of this function; everything contained within the opening and closing braces.
	Signature   string      `json:"signature"`    // The full definition of the function including receiver, name, arguments and return values.
	LineStart   int         `json:"line_start"`   // The line number in the associated source file where this function is initially defined.
	LineEnd     int         `json:"line_end"`     // The line number in the associated source file where the definition block ends.
	LineCount   int         `json:"line_count"`   // The total number of lines, including body, the interface occupies.
	astFunc     *ast.FuncDecl
	parent      *File
}

// NewFunction returns a function instance and attempts to populate all
// associated fields with meaningful values.
func NewFunction(fn *ast.FuncDecl, parent *File) *Function {
	return (&Function{
		Name:    fn.Name.Name,
		Comment: fn.Doc.Text(),
		astFunc: fn,
		parent:  parent,
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
// all comments immediately following the position of string "// Output:"
// belong to the output block.
func (f *Function) parseOutput() {
	var outputPos token.Pos
	if f.parent.astFile.Comments != nil {
		for _, cg := range f.parent.astFile.Comments {
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
	f.LineStart = f.parent.tokenSet.File(f.astFunc.Pos()).Line(f.astFunc.Pos())
	f.LineEnd = f.parent.tokenSet.Position(f.astFunc.End()).Line
	f.LineCount = f.LineEnd - f.LineStart
}

// parseBody attempts to make a few adjustments to the *ast.BlockStmt which
// represents the current function's body. We remove the opening and closing
// braces as well as the first occurrent `\t` sequence on each line.
func (f *Function) parseBody() {
	var b bytes.Buffer
	if err := format.Node(&b, f.parent.tokenSet, f.astFunc.Body); err != nil {
		log.Println(err)
	}

	f.Body = AdjustSource(fmt.Sprintf("%s", b.Bytes()), true)
}

// parseSignature attempts to determine the current function's type and assigns
// it to the Signature field of struct Function.
func (f *Function) parseSignature() {
	line := strings.TrimSpace(string(GetLinesFromFile(f.parent.Path, f.LineStart, f.LineStart)))
	f.Signature = line[:len(line)-1]
}

// String implements the Stringer inteface and returns the current function's
// signature.
func (f *Function) String() string {
	return f.Signature
}
