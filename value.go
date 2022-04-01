package gadget

import (
	"go/ast"
	"strings"
)

// Value represents a declared value in go; var, const, etc...
type Value struct {
	Kind     string `json:"kind"`
	Name     string `json:"name"`
	Line     int    `json:"line"`
	Body     string `json:"body"`
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
func (g *Value) Parse() *Value {
	g.Line = g.parent.tokenSet.File(g.astIdent.Pos()).Line(g.astIdent.Pos())
	g.Body = strings.TrimSpace(string(GetLinesFromFile(g.parent.Path, g.Line, g.Line)))

	return g
}

// String implements the Stringer inteface and returns the current values's
// body.
func (g *Value) String() string {
	return g.Body
}
