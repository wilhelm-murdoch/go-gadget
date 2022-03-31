package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wilhelm-murdoch/go-collection"
	"github.com/wilhelm-murdoch/go-gadget"
)

func main() {
	flagSource := flag.String("source", "*.go", "The directory to search for *.go files.")
	flag.Parse()

	packages := collection.New[*gadget.Package]()

	for _, path := range gadget.WalkGoFiles(flagSource) {
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

	// packages.Each(func(i int, p *gadget.Package) bool {
	// 	fmt.Println("package:", p.Name)
	// 	p.Files.Each(func(i int, f *gadget.File) bool {
	// 		fmt.Println("- file:", f.Name)
	// 		f.Functions.Each(func(i int, f *gadget.Function) bool {
	// 			fmt.Println("-- function:", f)
	// 			return false
	// 		})
	// 		f.Values.Each(func(i int, g *gadget.Value) bool {
	// 			fmt.Println("-- value:", g)
	// 			return false
	// 		})
	// 		f.Types.Each(func(i int, t *gadget.Type) bool {
	// 			fmt.Println("-- type:", t)
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
