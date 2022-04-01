package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wilhelm-murdoch/go-collection"
	"github.com/wilhelm-murdoch/go-gadget"
)

var (
	Version = "v0.0.1"
	Release = "development"
	Sha     = "xxxxxxxx"
)

func main() {
	var (
		flagSource = flag.String("source", "*.go", "The directory to search for *.go files.")
		// flgTemplate = flag.String("template", "README.tpl", "The path to the template you would like to evaluate.")
		// flgFormat   = flag.String("format", "json", "Chosen output format; json, template or debug.")
		flgVersion = flag.Bool("version", false, "Current version of gadget.")
	)
	flag.Parse()

	if *flgVersion {
		fmt.Printf("Version: %s, Release: %s, Sha: %s\n", Version, Release, Sha)
		os.Exit(0)
	}

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

	packages.Each(func(i int, p *gadget.Package) bool {
		p.Files.Each(func(i int, f *gadget.File) bool {
			f.Interfaces.Each(func(i int, t *gadget.Interface) bool {
				fmt.Println("-- interface:", t)
				return false
			})
			f.Structs.Each(func(i int, t *gadget.Struct) bool {
				// fmt.Println("-- struct:", t)
				return false
			})
			return false
		})
		return false
	})

	// switch *flgFormat {
	// case "debug":
	// 	spew.Dump(packages)
	// 	os.Exit(0)
	// case "template":
	// 	tpl := template.Must(
	// 		template.New(*flgTemplate).Funcs(sprig.FuncMap()).ParseFiles(*flgTemplate),
	// 	)

	// 	var buffer strings.Builder
	// 	if err := tpl.Execute(&buffer, packages); err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Print(html.UnescapeString(buffer.String()))
	// default:
	// 	fallthrough
	// case "json":
	// 	encoder := json.NewEncoder(os.Stdout)

	// 	if err := encoder.Encode(packages); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
}
