package gadget

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"

	"github.com/wilhelm-murdoch/go-collection"
)

// File represents a single file containing golang code.
type File struct {
	Name          string                             `json:"name"`           // The basename of the file.
	Path          string                             `json:"path"`           // The full path to the file as specified by the caller.
	Package       string                             `json:"package"`        // The name of the golang package associated with this file.
	IsMain        bool                               `json:"is_main"`        // Determines whether this file is part of package main.
	IsTest        bool                               `json:"is_test"`        // Determines whether this file is for golang tests.
	HasTests      bool                               `json:"has_tests"`      // Determines whether this file contains golang tests.
	HasBenchmarks bool                               `json:"has_benchmarks"` // Determines whether this file contains benchmark tests.
	HasExamples   bool                               `json:"has_examples"`   // Determines whether this file contains example tests.
	Imports       []string                           `json:"imports"`        // A list of strings containing all the current file's package imports.
	Values        *collection.Collection[*Value]     `json:"values"`         // A collection of declared golang values.
	Functions     *collection.Collection[*Function]  `json:"functions"`      // A collection of declared golang functions.
	Interfaces    *collection.Collection[*Interface] `json:"interfaces"`     // A collection of declared golang interfaces.
	Structs       *collection.Collection[*Struct]    `json:"structs"`        // A collection of declared golang structs.
	astFile       *ast.File
	tokenSet      *token.FileSet
}

// NewFile returns a file instance representing a file within a golang package.
// This function creates a new token set and parser instance representing the
// new file's abstract syntax tree ( AST ).
func NewFile(path string) (*File, error) {
	tokenSet := token.NewFileSet()
	astFile, err := parser.ParseFile(tokenSet, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	return (&File{
		Name:       filepath.Base(path),
		Path:       path,
		astFile:    astFile,
		tokenSet:   tokenSet,
		Values:     collection.New[*Value](),
		Functions:  collection.New[*Function](),
		Interfaces: collection.New[*Interface](),
		Structs:    collection.New[*Struct](),
	}).Parse(), nil
}

// Parse is responsible for walking through the current file's abstract syntax
// tree in order to populate it's fields. This includes imports, defined
// functions and methods, structs and interfaces and other declared values.
// ( Chainable )
func (f *File) Parse() *File {
	if strings.HasSuffix(f.Name, "_test.go") {
		f.IsTest = true
	}

	f.parsePackage()
	f.parseImports()
	f.parseFunctions()
	f.parseInterfaces()
	f.parseStructs()
	f.parseValues()

	return f
}

// parsePackage updates the current file with package-related data.
func (f *File) parsePackage() {
	f.Package = f.astFile.Name.Name
	if f.Package == "main" {
		f.IsMain = true
	}
}

// parseImports is responsible for creating a list of package imports that have
// been defined within the current file and assinging them to the appropriate
// Imports field.
func (f *File) parseImports() {
	for _, imp := range f.astFile.Imports {
		f.Imports = append(f.Imports, strings.ReplaceAll(imp.Path.Value, "\"", ""))
	}
}

// parseFunctions is responsible for creating abstract representations of
// functions and methods defined within the current file. All functions are
// added to the Functions collection.
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

			f.Functions.Push(NewFunction(fn, f))
		}
		return true
	})
}

// parseInterfaces is responsible for creating abstract representations of
// interfaces defined within the current file. All interfaces are added to the
// Interfaces collection.
func (f *File) parseInterfaces() {
	f.walk(func(node ast.Node) bool {
		ts, ok := node.(*ast.TypeSpec)
		if ok {
			if iface, ok := ts.Type.(*ast.InterfaceType); ok {
				f.Interfaces.Push(NewInterface(iface, ts, f))
			}
		}

		return true
	})
}

// parseInterfaces is responsible for creating abstract representations of
// structs defined within the current file. All interfaces are added to the
// Structs collection.
func (f *File) parseStructs() {
	f.walk(func(node ast.Node) bool {
		ts, ok := node.(*ast.TypeSpec)
		if ok {
			if st, ok := ts.Type.(*ast.StructType); ok {
				f.Structs.Push(NewStruct(st, ts, f))
			}
		}

		return true
	})
}

// parseValues is responsible for creating abstract representations of various
// general values such as const and var blocks. All values are added to the
// Values collection.
func (f *File) parseValues() {
	f.walk(func(node ast.Node) bool {
		switch gn := node.(type) {
		case *ast.ValueSpec:
			for _, ident := range gn.Names {
				f.Values.Push(NewValue(ident, f))
			}
		}
		return true
	})
}

// walk implements the walk interface which is used to step through syntax
// trees via a caller-supplied callback.
func (f *File) walk(fn func(ast.Node) bool) {
	ast.Walk(walker(fn), f.astFile)
}

// String implements the Stringer interface and returns the current files's
// path.
func (f *File) String() string {
	return f.Path
}
