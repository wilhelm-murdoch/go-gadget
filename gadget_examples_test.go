package gadget_test

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/wilhelm-murdoch/go-gadget"
)

func ExampleNewFile_functions() {
	if file, err := gadget.NewFile("sink/sink.go"); err == nil {
		file.Functions.Each(func(i int, function *gadget.Function) bool {
			fmt.Printf("%s defined between lines %d and %d\n", function.Name, function.LineStart, function.LineEnd)
			return false
		})
	}

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
}

func ExampleNewFile_structs() {
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
	// - &{1981 EmbeddedStructTest} on line 68
	// GenericStructTest is a struct with 4 fields:
	// - First on line 104
	// - Last on line 105
	// - Age on line 106
	// - private on line 107
}

func ExampleNewFile_json() {
	var buffer strings.Builder
	if file, err := gadget.NewFile("sink/sink.go"); err == nil {
		encoder := json.NewEncoder(&buffer)
		if err := encoder.Encode(file.Values.Items()); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(buffer.String())

	// Output:
	// [{"kind":"const","name":"ONE","line":9,"body":"ONE   = 1 // represents the number 1"},{"kind":"const","name":"TWO","line":10,"body":"TWO   = 2 // represents the number 2"},{"kind":"const","name":"THREE","line":11,"body":"THREE = 3 // represents the number 3"},{"kind":"var","name":"one","line":16,"body":"one   = \"one\"   // represents the english spelling of 1"},{"kind":"var","name":"two","line":17,"body":"two   = \"two\"   // represents the english spelling of 2"},{"kind":"var","name":"three","line":18,"body":"three = \"three\" // represents the english spelling of 3"},{"kind":"var","name":"collection","line":27,"body":"var collection map[string]map[string]string // this should be picked up as an inline comment."}]
}
