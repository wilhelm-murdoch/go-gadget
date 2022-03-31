package gadget

import (
	"fmt"
	"go/ast"
	"go/token"
)

type Type struct {
	Name      string
	Type      string
	LineStart int `json:"line_start"`
	LineEnd   int `json:"line_end"`
	LineCount int `json:"line_count"`
	Fields    []*TypeField
	astType   *ast.TypeSpec
	tokenSet  *token.FileSet
	astFile   *ast.File
}

type TypeField struct {
	Name       string `json:"name"`
	IsExported bool   `json:"is_exported"`
	Line       int    `json:"line"`
	Body       string `json:"body"`
}

func NewType(tp *ast.TypeSpec, ts *token.FileSet, f *ast.File) *Type {
	return (&Type{
		astType:  tp,
		tokenSet: ts,
		astFile:  f,
	}).Parse()
}

func (t *Type) Parse() *Type {
	var structType string
	var fields []TypeField
	switch st := t.astType.Type.(type) {
	case *ast.StructType:
		fmt.Println(t.astType.Name.Name, st.Pos(), "struct")
		for _, f := range st.Fields.List {
			for _, m := range f.Names {
				structType = fmt.Sprint(token.STRUCT)
				// t.Fields = append(t.Fields, &TypeField{m.Name, m.IsExported()})
				fmt.Println(m.Name, st.Pos(), "member of ", t.astType.Name.Name, m.IsExported())
			}
		}
	case *ast.InterfaceType:
		fmt.Println(t.astType.Name.Name, st.Pos(), "class")
		structType = fmt.Sprint(token.INTERFACE)
		for _, f := range st.Methods.List {
			for _, m := range f.Names {
				fmt.Println(m.Name, st.Pos(), "member of ", t.astType.Name.Name, m.IsExported())
			}
		}
	}
	fmt.Println(structType)
	return t
}

// String implements the Stringer inteface and returns the current package's
// name.
func (t *Type) String() string {
	return t.Name
}
