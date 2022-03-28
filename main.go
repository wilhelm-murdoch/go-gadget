package main

// template mode in + out
// pipe from stdin?
// logging
// add support for:
// - imports
// - structs
// - variables / constants

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/wilhelm-murdoch/go-collection"
)

func main() {
	var flgPath, flgIn, flgOut, flgFormat string

	flag.StringVar(&flgPath, "path", "*.go", "The directory to search for *.go files.")
	flag.StringVar(&flgIn, "in", "README.tpl", "The path to the template you would like to evaluate.")
	flag.StringVar(&flgOut, "out", "README.md", "The path to store the processed output.")
	flag.StringVar(&flgFormat, "format", "json", "Chosen output format; json, template or debug")
	flag.Parse()

	var files []string

	functions := collection.New[*Function]()
	packages := collection.New[string]()

	pattern, err := regexp.Compile(".+\\.go$")
	if err != nil {
		log.Fatal(err)
	}

	filepath.WalkDir(flgPath, func(path string, dir fs.DirEntry, err error) error {
		if err == nil && pattern.MatchString(dir.Name()) {
			files = append(files, path)
		}

		return nil
	})

	for _, file := range files {
		tfs := token.NewFileSet()
		parsed, err := parser.ParseFile(tfs, file, nil, parser.ParseComments)
		if err != nil {
			log.Println(err)
			break
		}

		packages.PushDistinct(parsed.Name.String())

		ast.Inspect(parsed, func(node ast.Node) bool {
			switch fn := node.(type) {
			case *ast.FuncDecl:
				var fntype bytes.Buffer
				if err := format.Node(&fntype, tfs, fn.Type); err != nil {
					panic(err)
				}

				var fnbody bytes.Buffer
				if err := format.Node(&fnbody, tfs, fn.Body); err != nil {
					panic(err)
				}

				functions.Push(&Function{
					Package:  parsed.Name.String(),
					FileName: filepath.Base(file),
					FilePath: file,
					Type:     fmt.Sprintf("%s", fntype.Bytes()),
					Name:     fn.Name.Name,
					Comment:  strings.ReplaceAll(fn.Doc.Text(), "\n", ""),
					Body:     fmt.Sprintf("%s", fnbody.Bytes()),
					Exported: fn.Name.IsExported(),
				})
			}

			return true
		})
	}

	// This is a bit janky at the moment. What I want is to associate an function
	// with its "example" counterpart, if it exists. This is so I can print out
	// example text along with the rest of the function's documentation. This may
	// grow a bit out-of-hand as support for other data types grows.
	functions.Each(func(i int, f1 *Function) bool {
		if strings.HasPrefix(f1.Name, "Example") {
			cmp := strings.Split(strings.Replace(f1.Name, "Example", "", 1), "_")

			found := functions.Find(func(i int, f2 *Function) bool {
				return f2.Package == strings.ToLower(cmp[0]) && f2.Name == cmp[1]
			})

			if found != nil {
				found.SetExample(f1.Body)
			}
		}
		return false
	})

	// This is the final data structure containing all of our collated
	// declarations. We munge it all together before finally passing it on to
	// whichever chosen output format.
	output := make(map[string]any, 0)

	output["functions"] = functions.Items()
	output["packages"] = packages.Items()

	switch flgFormat {
	case "debug":
		spew.Dump(output)
	case "template":
		tpl, err := template.ParseFiles(flgIn)
		if err != nil {
			panic(err)
		}

		var buffer strings.Builder
		if err = tpl.Execute(&buffer, output); err != nil {
			panic(err)
		}

		fmt.Print(html.UnescapeString(buffer.String()))
	default:
	case "json":
		encoder := json.NewEncoder(os.Stdout)

		if err := encoder.Encode(output); err != nil {
			panic(err)
		}
	}
}
