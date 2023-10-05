package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"html"
	"html/template"
	"os"
	"strings"
	"unicode"

	"github.com/Masterminds/sprig/v3"
	"github.com/urfave/cli/v2"
	"github.com/wilhelm-murdoch/go-collection"
	"github.com/wilhelm-murdoch/go-gadget"
)

// actionRootHandler is responsible for running the root command for the cli.
func actionRootHandler(c *cli.Context) error {
	// Initialise the collection responsible for collating any discovered top-
	// level go packages.
	packages := collection.New[*gadget.Package]()

	// Recursively walk the directory tree as specified by the `--source` flag. If
	// the flag points to a valid go file, begin parsing that instead.
	files := gadget.WalkGoFiles(c.String("source"))
	if len(files) == 0 {
		return errors.New("could not find any matching files ending in *.go")
	}

	// Start iterating through any discovered go files and start parsing.
	for _, path := range files {
		f, err := gadget.NewFile(path)
		if err != nil {
			return err
		}

		// Upsert any newly-discovered packages into the packages collection. Add
		// the current file to the package's `Files` sub-collection and move on.
		p := packages.Find(func(i int, p *gadget.Package) bool {
			return p.Name == f.Package
		})
		if p != nil {
			p.Files.Push(f)
			continue
		}

		// This is a new package, so let's add it to the packages collection and
		// assign the current file to the `Files` sub-collection and move on.
		p = gadget.NewPackage(f.Package)
		p.Files.Push(f)
		packages.Push(p)
	}

	// Iterate through all packages and their associated files and collect a list
	// of example functions for mapping.
	var examples []*gadget.Function
	packages.Each(func(i int, p *gadget.Package) bool {
		// Filter for files that test files and contian example functions:
		files := p.Files.Filter(func(f *gadget.File) bool {
			return f.IsTest && f.HasExamples
		})

		// Extract all example functions from the test files:
		files.Each(func(i int, f *gadget.File) bool {
			found := f.Functions.Filter(func(f *gadget.Function) bool { return f.IsExample })
			examples = append(examples, found.Items()...)
			return false
		})
		return false
	})

	// Now, we must again iterate through all packages and their files to map each
	// example function to their associated function / method.
	packages.Each(func(i int, p *gadget.Package) bool {
		// Iterate through all known examples and attempt to map them to their
		// corresponding functions / methods in the main collection.
		for _, example := range examples {
			parts := strings.Split(strings.TrimPrefix(example.Name, "Example"), "_")

			// If we have at least 2 parts and the last part begins with a lowercase
			// character, we can safely assume this is a part of a set of examples
			// targeting a single function / method.
			if len(parts) >= 2 && unicode.IsLower([]rune(parts[len(parts)-1])[0]) {
				parts = parts[:len(parts)-1]
			}

			// If we have 2 parts left, we can assume with have a target type and
			// target method. If we a single part left, it's going to be for a
			// standalone function.
			var predicate func(int, *gadget.Function) bool
			if len(parts) == 2 {
				predicate = func(i int, f *gadget.Function) bool {
					return f.Receiver == parts[0] && f.Name == parts[1]
				}
			} else if len(parts) == 1 {
				predicate = func(i int, f *gadget.Function) bool {
					return f.Name == parts[0]
				}
			}

			// Use the appropriate predicate function and push the example function
			// into the associated function's collection of examples.
			p.Files.Each(func(i int, f *gadget.File) bool {
				if found := f.Functions.Find(predicate); found != nil {
					found.Examples.Push(example)
					return true
				}
				return false
			})
		}
		return false
	})

	// This cli supports three output modes: debug ( using ast.Print ), JSON and
	// template support using go's `html/template` package in addition to sprig's
	// set of template functions using `sprig.FuncMap()`.
	switch c.String("format") {
	default:
		fallthrough
	case "json":
		encoder := json.NewEncoder(os.Stdout)
		if err := encoder.Encode(packages.Items()); err != nil {
			return err
		}
	case "debug":
		packages.Each(func(i int, p *gadget.Package) bool {
			p.Files.Each(func(i int, f *gadget.File) bool {
				astFile, tokenSet := f.GetAstAttributes()
				ast.Print(tokenSet, astFile)
				return false
			})
			return false
		})
	case "template":
		tpl, err := template.New(c.String("template")).Funcs(sprig.FuncMap()).ParseFiles(c.String("template"))
		if err != nil {
			return err
		}

		var buffer strings.Builder
		if err := tpl.Execute(&buffer, packages.Items()); err != nil {
			return err
		}

		fmt.Println(html.UnescapeString(buffer.String()))
	}

	return nil
}
