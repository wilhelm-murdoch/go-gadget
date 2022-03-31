package gadget

import (
	"go/ast"
	"go/token"
)

type Type struct {
	Name     string
	Fields   []*TypeField
	astType  *ast.TypeSpec
	tokenSet *token.FileSet
	astFile  *ast.File
}

type TypeField struct {
}

func NewType(tp *ast.TypeSpec, ts *token.FileSet, f *ast.File) *Type {
	return (&Type{
		astType:  tp,
		tokenSet: ts,
		astFile:  f,
	}).Parse()
}

func (t *Type) Parse() *Type {
	// switch st := spec.Type.(type) {
	// case *ast.StructType:
	// 	fmt.Println(spec.Name.Name, st.Pos(), "struct")
	// 	for _, f := range st.Fields.List {
	// 		for _, m := range f.Names {
	// 			fmt.Println(m.Name, st.Pos(), "member of ", spec.Name.Name)
	// 		}
	// 	}
	// case *ast.InterfaceType:
	// 	fmt.Println(spec.Name.Name, st.Pos(), "class")
	// 	for _, f := range st.Methods.List {
	// 		for _, m := range f.Names {
	// 			fmt.Println(m.Name, m.Pos(), "func of ", spec.Name.Name)
	// 		}
	// 	}
	// default:
	// 	fmt.Println(spec.Name.Name, st.Pos(), "type")
	// }

	return t
}
