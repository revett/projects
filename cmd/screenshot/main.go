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
		Flags: flag.Global(),
		Commands: []*cli.Command{
			cmd.Element(),
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
