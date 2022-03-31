package gadget

import (
	"go/ast"
	"go/token"

	"github.com/wilhelm-murdoch/go-collection"
)

type Interface struct {
	Name      string `json:"name"`
	LineStart int    `json:"line_start"`
	LineEnd   int    `json:"line_end"`
	LineCount int    `json:"line_count"`
	Fields    *collection.Collection[*Field]
	astType   *ast.InterfaceType
	astSpec   *ast.TypeSpec
	tokenSet  *token.FileSet
	astFile   *ast.File
}

func NewInterface(it *ast.InterfaceType, ts *ast.TypeSpec, fs *token.FileSet, f *ast.File) *Interface {
	return (&Interface{
		Fields:   collection.New[*Field](),
		astType:  it,
		astSpec:  ts,
		tokenSet: fs,
		astFile:  f,
	}).Parse()
}

func (i *Interface) Parse() *Interface {
	if i.astSpec != nil {
		i.Name = i.astSpec.Name.Name
	}
	return i
}

// String implements the Stringer inteface and returns the current package's
// name.
func (i *Interface) String() string {
	return i.Name
}
