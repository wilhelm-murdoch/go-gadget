package gadget

import (
	"go/ast"
	"strings"
)

// Value represents a declared value in go; var, const, etc...
type Value struct {
	Kind     string `json:"kind"` // Describes the current value's type, eg; CONST or VAR.
	Name     string `json:"name"` // The name of the value.
	Line     int    `json:"line"` // The line number within the associated source file in which this value was originally defined.
	Body     string `json:"body"` // The full content of the associated statement.
	astIdent *ast.Ident
	parent   *File
}

// NewValue returns a Value instance.
func NewValue(id *ast.Ident, parent *File) *Value {
	return (&Value{
		Name:     id.String(),
		Kind:     id.Obj.Kind.String(),
		astIdent: id,
		parent:   parent,
	}).Parse()
}

// Parse is responsible for browsing through f.astIdent and f.tokenSet to
// populate the current value's fields. ( Chainable )
func (v *Value) Parse() *Value {
	v.Line = v.parent.tokenSet.File(v.astIdent.Pos()).Line(v.astIdent.Pos())
	v.Body = strings.TrimSpace(string(GetLinesFromFile(v.parent.Path, v.Line, v.Line)))

	return v
}

// String implements the Stringer interface and returns the current values's
// body.
func (v *Value) String() string {
	return v.Body
}
