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

// NewFile
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

// Parse
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

			f.Functions.Push(NewFunction(fn, f))
		}
		return true
	})
}

// Foo is a stupid interface. Just look at this fucking idiot!
type Foo interface {
	// I am on top of the blup
	Blup(say string) string // this is for all the blups out there
	// I am up here
	Boop() int // fuck the boops
}

type Mammal interface {
	getType() string
	canFly() bool
	feed(foodCount int) string
}

// parseInterfaces
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

// parseStructs
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

// parseValues
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

// walk
func (f *File) walk(fn func(ast.Node) bool) {
	ast.Walk(walker(fn), f.astFile)
}
