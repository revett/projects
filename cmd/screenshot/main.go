package main

import (
	"log"
	"os"
	"sort"

	"github.com/revett/projects/internal/screenshot/cmd"
	"github.com/revett/projects/internal/screenshot/flag"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Authors: []*cli.Author{
			{
				Email: "@revcd",
				Name:  "Charlie Revett",
			},
		},
		Commands: []*cli.Command{
			cmd.Element(),
		},
		Flags: flag.Global(),
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
