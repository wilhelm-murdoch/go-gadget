package gadget

import (
	"go/ast"
	"go/token"
	"strings"
)

// Value represents a declared value in go.
type Value struct {
	Kind     string `json:"kind"`
	Name     string `json:"name"`
	Line     int    `json:"line"`
	Body     string `json:"body"`
	astIdent *ast.Ident
	tokenSet *token.FileSet
	parent   *File
}

// NewValue returns a Value instance.
func NewValue(id *ast.Ident, ts *token.FileSet, parent *File) *Value {
	return (&Value{
		astIdent: id,
		tokenSet: ts,
		parent:   parent,
	}).Parse()
}

// Parse is responsible for browsing through f.astIdent and f.tokenSet to
// populate the current value's fields. ( Chainable )
func (g *Value) Parse() *Value {
	g.Kind = g.astIdent.Obj.Kind.String()
	g.Name = g.astIdent.String()
	g.Line = g.tokenSet.File(g.astIdent.Pos()).Line(g.astIdent.Pos())
	g.Body = strings.TrimSpace(string(GetLinesFromFile(g.parent.Path, g.Line, g.Line)))

	return g
}

// String implements the Stringer inteface and returns the current values's
// body.
func (g *Value) String() string {
	return g.Body
}
