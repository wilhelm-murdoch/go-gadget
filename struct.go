package gadget

import (
	"go/ast"

	"github.com/wilhelm-murdoch/go-collection"
)

type Struct struct {
	Name      string `json:"name"`
	LineStart int    `json:"line_start"`
	LineEnd   int    `json:"line_end"`
	LineCount int    `json:"line_count"`
	Fields    *collection.Collection[*Field]
	astType   *ast.StructType
	astSpec   *ast.TypeSpec
	parent    *File
}

func NewStruct(st *ast.StructType, ts *ast.TypeSpec, parent *File) *Struct {
	return (&Struct{
		Fields:  collection.New[*Field](),
		astType: st,
		astSpec: ts,
		parent:  parent,
	}).Parse()
}

func (s *Struct) Parse() *Struct {
	s.Name = s.astSpec.Name.Name
	return s
}

// String implements the Stringer inteface and returns the current package's
// name.
func (s *Struct) String() string {
	return s.Name
}
