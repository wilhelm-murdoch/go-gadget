package main

// template mode in + out
// default output is json

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	ignoredPatterns = []string{"_benchmark.go", "_test.go"}
	manifest        map[string]*Package
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
	var path, in, out, format string

	flag.StringVar(&path, "path", "*.go", "The directory to search for *.go files.")
	flag.StringVar(&in, "in", "README.tpl", "The path to the template you would like to evaluate.")
	flag.StringVar(&out, "out", "README.md", "The path to store the processed output.")
	flag.StringVar(&format, "format", "json", "The desired output format; JSON support only.")
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
		parsed, err := parser.ParseFile(token.NewFileSet(), file, nil, parser.ParseComments)
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

		ast.Inspect(parsed, func(n ast.Node) bool {
			switch fn := n.(type) {
			case *ast.FuncDecl:
				fnc := &Function{
					Name:    fn.Name.Name,
					Comment: fn.Doc.Text(),
				}

				f.Functions.Push(fnc)
			}

			return true
		})

		pkg.Files.Push(f)
	}

	t, err := template.ParseFiles(in)

	if err != nil {
		panic(err)
	}

	packages := []*Package{}

	manifest.Each(func(i int, p *Package) bool {
		packages = append(packages, p)
		return false
	})

	err = t.Execute(os.Stdout, packages)

	if err != nil {
		panic(err)
	}

	// manifest.Each(func(i int, p *Package) bool {
	// 	fmt.Println("Package: ", p.Name)
	// 	p.Files.Each(func(i int, f *File) bool {
	// 		fmt.Println("\tFile: ", f.Name)
	// 		f.Functions.Each(func(i int, f *Function) bool {
	// 			fmt.Println("\t\tFunc", f.Name)
	// 			return false
	// 		})
	// 		return false
	// 	})
	// 	return false
	// })
}
