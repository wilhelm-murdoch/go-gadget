package gadget

import (
	"go/ast"
	"go/token"

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
	tokenSet  *token.FileSet
	astFile   *ast.File
}

func NewStruct(st *ast.StructType, ts *ast.TypeSpec, fs *token.FileSet, f *ast.File) *Struct {
	return (&Struct{
		Fields:   collection.New[*Field](),
		astType:  st,
		astSpec:  ts,
		tokenSet: fs,
		astFile:  f,
	}).Parse()
}

func (s *Struct) Parse() *Struct {
	return s
}

// String implements the Stringer inteface and returns the current package's
// name.
func (s *Struct) String() string {
	return s.Name
}
