package gadget

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"

	"github.com/wilhelm-murdoch/go-collection"
)

// Struct File represents a file containing Go code.
type File struct {
	Name          string                            `json:"name"`
	Path          string                            `json:"path"`
	Package       string                            `json:"package"`
	IsMain        bool                              `json:"is_main"`
	IsTest        bool                              `json:"is_test"`
	HasTests      bool                              `json:"has_tests"`
	HasBenchmarks bool                              `json:"has_benchmarks"`
	HasExamples   bool                              `json:"has_examples"`
	Imports       []string                          `json:"imports"`
	General       *collection.Collection[*General]  `json:"general"`
	Functions     *collection.Collection[*Function] `json:"functions"`
	Types         *collection.Collection[*Type]     `json:"types"`
	astFile       *ast.File
	tokenSet      *token.FileSet
}

// NewFile
func NewFile(path string) (*File, error) {
	tokenSet := token.NewFileSet()
	astFile, err := parser.ParseFile(tokenSet, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	return (&File{
		Name:      filepath.Base(path),
		Path:      path,
		astFile:   astFile,
		tokenSet:  tokenSet,
		General:   collection.New[*General](),
		Functions: collection.New[*Function](),
		Types:     collection.New[*Type](),
	}).Parse(), nil
}

// Parse
func (f *File) Parse() *File {
	if strings.HasSuffix(f.Name, "_test.go") {
		f.IsTest = true
	}

	f.parsePackage()
	f.parseImports()
	f.parseFunctions()
	f.parseTypes()
	f.parseGeneral()

	return f
}

// parsePackage
func (f *File) parsePackage() {
	f.Package = f.astFile.Name.Name
	if f.Package == "main" {
		f.IsMain = true
	}
}

// parseImports
func (f *File) parseImports() {
	for _, imp := range f.astFile.Imports {
		f.Imports = append(f.Imports, imp.Path.Value)
	}
}

// parseFunctions
func (f *File) parseFunctions() {
	f.walk(func(node ast.Node) bool {
		switch fn := node.(type) {
		case *ast.FuncDecl:
			if strings.HasPrefix(fn.Name.Name, "Example") {
				f.HasExamples = true
			} else if strings.HasPrefix(fn.Name.Name, "Benchmark") {
				f.HasBenchmarks = true
			} else if strings.HasPrefix(fn.Name.Name, "Test") {
				f.HasTests = true
			}

			f.Functions.Push(NewFunction(fn, f.tokenSet, f.astFile))
		}
		return true
	})
}

// parseTypes
func (f *File) parseTypes() {
	f.walk(func(node ast.Node) bool {
		switch tp := node.(type) {
		case *ast.TypeSpec:
			f.Types.Push(NewType(tp, f.tokenSet, f.astFile))
		}

		return true
	})
}

// parseGeneral
func (f *File) parseGeneral() {
	f.walk(func(node ast.Node) bool {
		switch gn := node.(type) {
		case *ast.ValueSpec:
			for _, id := range gn.Names {
				fmt.Println(id.Obj.Kind, id.String(), id.Obj.Decl, f.tokenSet.File(f.astFile.Pos()).Line(id.Pos()))
			}
		}
		return true
	})
}

// walk
func (f *File) walk(fn func(ast.Node) bool) {
	ast.Walk(walker(fn), f.astFile)
}

// walker
type walker func(ast.Node) bool

// Visit
func (w walker) Visit(node ast.Node) ast.Visitor {
	if w(node) {
		return w
	}
	return nil
}
