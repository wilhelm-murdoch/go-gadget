package main

// template mode in + out
// default output is json
// pipe from stdin?

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	// ignoredPatterns = []string{"_benchmark.go", "_test.go"}
	ignoredPatterns = []string{"_benchmark.go"}
	manifest        = make(map[string]*Package)
)

func endsWithIgnoredPattern(s string) bool {
	for _, pattern := range ignoredPatterns {
		if strings.HasSuffix(s, pattern) {
			return true
		}
	}
	return false
}

func main() {
	var path, in, out string

	flag.StringVar(&path, "path", "*.go", "The directory to search for *.go files.")
	flag.StringVar(&in, "in", "README.tpl", "The path to the template you would like to evaluate.")
	flag.StringVar(&out, "out", "README.md", "The path to store the processed output.")
	flag.Parse()

	var files []string

	pattern, err := regexp.Compile(".+\\.go$")
	if err != nil {
		log.Fatal(err)
	}

	filepath.WalkDir(path, func(path string, dir fs.DirEntry, err error) error {
		if err == nil && pattern.MatchString(dir.Name()) && !endsWithIgnoredPattern(dir.Name()) {
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

		pkg := parsed.Name.String()

		_, ok := manifest[pkg]
		if !ok {
			manifest[pkg] = &Package{Name: pkg}
		}

		f := &File{
			Name: filepath.Base(file),
			Path: file,
		}

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

				f.Functions = append(f.Functions, &Function{
					Type:    fmt.Sprintf("%s", fntype.Bytes()),
					Name:    fn.Name.Name,
					Comment: fn.Doc.Text(),
					Body:    fmt.Sprintf("%s", fnbody.Bytes()),
				})
			}

			return true
		})

		manifest[pkg].Files = append(manifest[pkg].Files, f)
	}

	encoder := json.NewEncoder(os.Stdout)

	err = encoder.Encode(manifest)
	if err != nil {
		panic(err)
	}

	// template, err := template.ParseFiles(in)
	// if err != nil {
	// 	panic(err)
	// }

	// if err = template.Execute(os.Stdout, manifest); err != nil {
	// 	panic(err)
	// }
}
