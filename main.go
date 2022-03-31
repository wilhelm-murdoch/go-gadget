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

	packages := collection.New[*Package]()

	for _, path := range walkGoFiles(flagSource) {
		f, err := NewFile(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		p := packages.Find(func(i int, p *Package) bool {
			return p.Name == f.Package
		})
		if p != nil {
			p.Files.Push(f)
			continue
		}

		p = NewPackage(f.Package)
		p.Files.Push(f)
		packages.Push(p)
	}

	// packages.Each(func(i int, p *Package) bool {
	// 	fmt.Println(p.Name)
	// 	p.Files.Each(func(i int, f *File) bool {
	// 		fmt.Println("-", f.Name)
	// 		filtered := f.Functions.Filter(func(f *Function) bool {
	// 			return strings.HasPrefix(f.Name, "Example")
	// 		})
	// 		filtered.Each(func(i int, f *Function) bool {
	// 			fmt.Println("--", f.Body)
	// 			return false
	// 		})
	// 		return false
	// 	})
	// 	return false
	// })

	// encoder := json.NewEncoder(os.Stdout)

	// if err := encoder.Encode(packages.Items()); err != nil {
	// 	log.Fatal(err)
	// }
}
