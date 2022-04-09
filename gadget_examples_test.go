package gadget_test

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/wilhelm-murdoch/go-gadget"
)

func ExampleNewFile_functions() {
	if file, err := gadget.NewFile("function.go"); err == nil {
		file.Functions.Each(func(i int, function *gadget.Function) bool {
			fmt.Printf("%s defined between lines %d and %d\n", function.Name, function.LineStart, function.LineEnd)
			return false
		})
	}

	// Output:
	// NewFunction defined between lines 36 and 43
	// Parse defined between lines 47 and 68
	// parseReceiver defined between lines 71 and 88
	// parseOutput defined between lines 94 and 114
	// parseLines defined between lines 118 and 122
	// parseBody defined between lines 127 and 136
	// parseSignature defined between lines 140 and 143
	// String defined between lines 146 and 148
}

func ExampleNewFile_structs() {
	if file, err := gadget.NewFile("function.go"); err == nil {
		file.Types.Each(func(i int, t *gadget.Type) bool {
			fmt.Printf("%s is a %s with %d fields:\n", t.Name, t.Kind, t.Fields.Length())
			t.Fields.Each(func(i int, f *gadget.Field) bool {
				fmt.Printf("- %s on line %d\n", f.Name, f.Line)
				return false
			})
			return false
		})
	}

	// Output:
	// Function is a struct with 16 fields:
	// - Name on line 16
	// - IsTest on line 17
	// - IsBenchmark on line 18
	// - IsExample on line 19
	// - IsExported on line 20
	// - IsMethod on line 21
	// - Receiver on line 22
	// - Doc on line 23
	// - Output on line 24
	// - Body on line 25
	// - Signature on line 26
	// - LineStart on line 27
	// - LineEnd on line 28
	// - LineCount on line 29
	// - astFunc on line 30
	// - parent on line 31
}

func ExampleNewFile_json() {
	var buffer strings.Builder
	if file, err := gadget.NewFile("value.go"); err == nil {
		if t, ok := file.Types.AtFirst(); ok {
			encoder := json.NewEncoder(&buffer)
			if err := encoder.Encode(t.Fields.Items()); err != nil {
				fmt.Println(err)
			}
		}
	}

	fmt.Println(buffer.String())

	// Output:
	// [{"name":"Kind","is_exported":true,"is_embedded":false,"line":10,"body":"Kind string `json:\"kind\"`","comment":"Describes the current value's type, eg; CONST or VAR."},{"name":"Name","is_exported":true,"is_embedded":false,"line":11,"body":"Name string `json:\"name\"`","comment":"The name of the value."},{"name":"Line","is_exported":true,"is_embedded":false,"line":12,"body":"Line int `json:\"line\"`","comment":"The line number within the associated source file in which this value was originally defined."},{"name":"Body","is_exported":true,"is_embedded":false,"line":13,"body":"Body string `json:\"body\"`","comment":"The full content of the associated statement."},{"name":"astIdent","is_exported":false,"is_embedded":false,"line":14,"body":"astIdent *ast.Ident"},{"name":"parent","is_exported":false,"is_embedded":false,"line":15,"body":"parent *File"}]
}
