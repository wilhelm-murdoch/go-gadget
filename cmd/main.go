package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"html/template"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Masterminds/sprig"
	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"
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

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("Version: %s, Stage: %s, Commit: %s, Date: %s\n", Version, Stage, Commit, Date)
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Usage:   "print only the version",
		Aliases: []string{"v"},
	}

	app := &cli.App{
		Name:     "gadget",
		Usage:    "inspect your code via a small layer of abstraction over Go's AST package",
		Version:  Version,
		Compiled: time.Now(),
		Authors: []*cli.Author{{
			Name:  "Wilhelm Murdoch",
			Email: "wilhelm@devilmayco.de",
		}},
		Copyright: "(c) 2022 Wilhelm Codes ( https://wilhelm.codes )",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "source",
				Usage:   "path to the target go source file, or directory containing go source files.",
				Value:   ".",
				Aliases: []string{"s"},
			},
			&cli.StringFlag{
				Name:    "format",
				Usage:   "the output format of the results as json, template or debug.",
				Value:   "json",
				Aliases: []string{"f"},
			},
			&cli.StringFlag{
				Name:    "template",
				Usage:   "if the template format is selected, this is the path to the template file to use.",
				Value:   "README.tpl",
				Aliases: []string{"t"},
			},
		},
		Action: actionRootHandler,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func actionRootHandler(c *cli.Context) error {
	packages := collection.New[*gadget.Package]()

	files := gadget.WalkGoFiles(c.String("source"))
	if len(files) == 0 {
		return errors.New("could not find any matching files ending in *.go")
	}

	for _, path := range files {
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

	switch c.String("format") {
	case "debug":
		spew.Dump(packages)
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
	default:
		fallthrough
	case "json":
		encoder := json.NewEncoder(os.Stdout)
		if err := encoder.Encode(packages.Items()); err != nil {
			return err
		}
	}

	return nil
}
