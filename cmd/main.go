package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
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
