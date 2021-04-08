package cmd

import (
	"log"

	"github.com/revett/projects/internal/screenshot/browser"
	"github.com/revett/projects/internal/screenshot/imgio"
	"github.com/revett/projects/internal/screenshot/page"
	"github.com/tebeka/selenium"
	"github.com/urfave/cli/v2"
)

// Element is a CLI command that takes a screenshot of a specific web element.
func Element() *cli.Command {
	return &cli.Command{
		Name:    "element",
		Aliases: []string{"e"},
		Usage:   "take screenshot of a web element",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "selector",
				Aliases:  []string{"s"},
				Required: true,
				Usage:    "css selector for element",
			},
		},
		Action: elementAction,
	}
}

func elementAction(c *cli.Context) error {
	b, err := browser.New(
		selenium.NewRemote, c.String("browserName"), c.String("host"),
	)
	if err != nil {
		return err
	}

	defer func() {
		if err := b.Quit(); err != nil {
			log.Println("err")
		}
	}()

	om := page.New(b)

	err = om.Visit(c.String("url"))
	if err != nil {
		return err
	}

	err = om.WaitForElement(c.String("selector"), c.Int("timeout"))
	if err != nil {
		return err
	}

	bytes, err := om.ScreenshotElement(c.String("selector"))
	if err != nil {
		return err
	}

	return imgio.Write(bytes)
}
