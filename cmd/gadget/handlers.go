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

	"github.com/Masterminds/sprig"
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

	// This cli supports three output modes: debug ( using spew ), JSON and
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
