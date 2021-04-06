package cmd

import (
	"log"

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
		Action: func(c *cli.Context) error {
			log.Printf("setting webdriver browsername: %s", c.String("browserName"))
			caps := selenium.Capabilities{
				"browserName": c.String("browserName"),
			}

			log.Printf("connecting to local selenium host: %s", c.String("host"))
			wd, err := selenium.NewRemote(caps, c.String("host"))
			if err != nil {
				return err
			}

			defer func() {
				if err := wd.Quit(); err != nil {
					log.Println("err")
				}
			}()

			om := page.New(wd)
			err = om.Visit(c.String("url"))
			if err != nil {
				return err
			}

			err = om.WaitForElement(c.String("selector"), c.Int("timeout"))
			if err != nil {
				return err
			}

			b, err := om.ScreenshotElement(c.String("selector"))
			if err != nil {
				return err
			}

			return imgio.Write(b)
		},
	}
}
