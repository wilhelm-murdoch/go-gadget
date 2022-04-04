package gadget

import (
	"go/ast"
	"strings"

	"github.com/wilhelm-murdoch/go-collection"
)

const (
	KIND_INTERFACE   = "interface"   // Used if the type is an interface with methods.
	KIND_STRUCT      = "struct"      // Used if the type is a struct with fields.
	KIND_ARRAY       = "array"       // Used for array types.
	KIND_FUNC        = "function"    // Use function types.
	KIND_CHAN        = "channel"     // Used for channel types.
	KIND_MAP         = "map"         // Used for map types.
	KIND_UNSUPPORTED = "unsupported" // Used if no matching kinds can be found.
)

// Type represents a golang type definition.
type Type struct {
	Name      string                         `json:"name"`              // The name of the struct.
	Kind      string                         `json:"kind"`              // Determines the kind of type, eg; interface or struct.
	LineStart int                            `json:"line_start"`        // The line number in the associated source file where this struct is initially defined.
	LineEnd   int                            `json:"line_end"`          // The line number in the associated source file where the definition block ends.
	LineCount int                            `json:"line_count"`        // The total number of lines, including body, the struct occupies.
	Comment   string                         `json:"comment,omitempty"` // Any inline comments associated with the struct.
	Doc       string                         `json:"doc,omitempty"`     // The comment block directly above this struct's definition.
	Signature string                         `json:"signature"`         // The full definition of the struct itself.
	Body      string                         `json:"body,omitempty"`    // The full body of the struct sourced directly from the associated file; comments included.
	Fields    *collection.Collection[*Field] `json:"fields,omitempty"`  // A collection of fields and their associated metadata.
	astSpec   *ast.TypeSpec
	parent    *File
}

// NewType returns an struct instance and attempts to populate all associated
// fields with meaningful values.
func NewType(ts *ast.TypeSpec, parent *File) *Type {
	return (&Type{
		Name:    ts.Name.Name,
		Fields:  collection.New[*Field](),
		Doc:     ts.Doc.Text(),
		astSpec: ts,
		parent:  parent,
	}).Parse()
}

// Parse is responsible for browsing through f.astSpec, f.astType, f.parent to
// populate the current struct's fields. ( Chainable )
func (t *Type) Parse() *Type {
	t.parseLines()
	t.parseBody()
	t.parseSignature()
	t.parseFields()

	return t
}

// parseLines determines the current struct's opening and closing line
// positions.
func (t *Type) parseLines() {
	t.LineStart = t.parent.tokenSet.File(t.astSpec.Pos()).Line(t.astSpec.Pos())
	t.LineEnd = t.parent.tokenSet.Position(t.astSpec.End()).Line
	t.LineCount = (t.LineEnd + 1) - t.LineStart
}

// parseBody attempts to make a few adjustments to the *ast.BlockStmt which
// represents the current struct's body. We remove the opening and closing
// braces as well as the first occurrent `\t` sequence on each line.
func (t *Type) parseBody() {
	t.Body = AdjustSource(string(GetLinesFromFile(t.parent.Path, t.LineStart+1, t.LineEnd-1)), false)
}

// parseSignature attempts to determine the current structs's type and assigns
// it to the Signature field of struct Function.
func (t *Type) parseSignature() {
	line := strings.TrimSpace(string(GetLinesFromFile(t.parent.Path, t.LineStart, t.LineStart)))
	t.Signature = line[:len(line)-1]
}

// parseFields iterates through the struct's list of defined methods to
// populate the Fields collection.
func (t *Type) parseFields() {
	switch tp := t.astSpec.Type.(type) {
	case *ast.StructType:
		t.Kind = KIND_STRUCT
		for _, field := range tp.Fields.List {
			t.Fields.Push(NewField(field, t.parent))
		}
	case *ast.InterfaceType:
		t.Kind = KIND_INTERFACE
		for _, method := range tp.Methods.List {
			t.Fields.Push(NewField(method, t.parent))
		}
	case *ast.ArrayType:
		t.Kind = KIND_ARRAY
	case *ast.FuncType:
		t.Kind = KIND_FUNC
	case *ast.ChanType:
		t.Kind = KIND_CHAN
	case *ast.MapType:
		t.Kind = KIND_MAP
	default:
		t.Kind = KIND_UNSUPPORTED
	}
}

// String implements the Stringer struct and returns the current package's name.
func (t *Type) String() string {
	return t.Name
}
