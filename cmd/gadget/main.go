package main

// https://github.com/ArjenL/taggo/blob/b6906610871a22b53941432ce166423844c34b8b/main.go#L237

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/wilhelm-murdoch/go-collection"
	"github.com/wilhelm-murdoch/go-gadget"
)

var (
	this string // a comment for this

	// a comment for that
	that string
)

const (
	Boo  = "ghost"
	Hiss = "cat"
)

func walkGoFiles(path *string) []string {
	var files []string

	pattern, err := regexp.Compile(".+\\.go$")
	if err != nil {
		log.Fatal(err)
	}

	filepath.WalkDir(*path, func(path string, dir fs.DirEntry, err error) error {
		if err == nil && pattern.MatchString(dir.Name()) {
			files = append(files, path)
		}
		return nil
	})

	return files
}

func main() {
	var (
		flagSource = flag.String("source", "*.go", "The directory to search for *.go files.")
	)
	flag.Parse()

	packages := collection.New[*gadget.Package]()

	for _, path := range walkGoFiles(flagSource) {
		f, err := gadget.NewFile(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		p := packages.Find(func(i int, p *gadget.Package) bool {
			return p.Name == f.Package
		})
		if p != nil {
			p.Files.Push(f)
			continue
		}

		p = gadget.NewPackage(f.Package)
		p.Files.Push(f)
		packages.Push(p)
	}

	packages.Each(func(i int, p *gadget.Package) bool {
		fmt.Println("package:", p.Name)
		p.Files.Each(func(i int, f *gadget.File) bool {
			fmt.Println("- file:", f.Name)
			f.General.Each(func(i int, g *gadget.General) bool {
				fmt.Println("-- general:", g)
				return false
			})
			return false
		})
		return false
	})

	// encoder := json.NewEncoder(os.Stdout)

	// if err := encoder.Encode(packages.Items()); err != nil {
	// 	log.Fatal(err)
	// }
}
