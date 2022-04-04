package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/davecgh/go-spew/spew"
	"github.com/wilhelm-murdoch/go-collection"
	"github.com/wilhelm-murdoch/go-gadget"
)

const (
	// Version describes the version of the current build.
	Version = "dev"

	// Commit describes the commit of the current build.
	Commit = "none"

	// Date describes the date of the current build.
	Date = "unknown"

	// Release describes the stage of the current build, eg; development, production, etc...
	Stage = "unknown"
)

var (
	flagSource  = flag.String("source", "*.go", "The directory to search for *.go files.")
	flgTemplate = flag.String("template", "README.tpl", "The path to the template you would like to evaluate.")
	flgFormat   = flag.String("format", "json", "Chosen output format; json, template or debug.")
	flgVersion  = flag.Bool("version", false, "Current version of gadget.")
)

func main() {
	flag.Parse()

	if *flgVersion {
		fmt.Printf("Version: %s, Stage: %s, Commit: %s\n, Built On: %s", Version, Stage, Commit, Date)
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

	switch *flgFormat {
	case "debug":
		spew.Dump(packages)
		os.Exit(0)
	case "template":
		tpl := template.Must(
			template.New(*flgTemplate).Funcs(sprig.FuncMap()).ParseFiles(*flgTemplate),
		)

		var buffer strings.Builder
		if err := tpl.Execute(&buffer, packages.Items()); err != nil {
			log.Fatal(err)
		}

		fmt.Print(html.UnescapeString(buffer.String()))
	default:
		fallthrough
	case "json":
		encoder := json.NewEncoder(os.Stdout)

		if err := encoder.Encode(packages.Items()); err != nil {
			log.Fatal(err)
		}
	}
}
