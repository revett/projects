package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/revett/projects/internal/screenshot/page"
	"github.com/tebeka/selenium"
	"github.com/urfave/cli/v2"
)

// Element is a CLI command that takes a screenshot of a specific web element.
func Element() *cli.Command {
	return &cli.Command{
		Name:    "element",
		Aliases: []string{"e"},
		Usage:   "take screenshot of web element",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "selector",
				Aliases:  []string{"s"},
				Required: true,
				Usage:    "css selector for element",
			},
		},
		Action: func(c *cli.Context) error {
			log.Printf("setting webdriver browsername: %s", c.String("bn"))
			caps := selenium.Capabilities{
				"browserName": c.String("bn"),
			}

			log.Printf("connecting to local selenium host: %s", c.String("host"))
			wd, err := selenium.NewRemote(caps, c.String("host"))
			if err != nil {
				return err
			}
			defer wd.Quit() // nolint:errcheck

			log.Printf("opening webpage: %s", c.String("u"))
			if err := wd.Get(c.String("u")); err != nil {
				return err
			}

			log.Printf(
				"waiting for element (up to %ds): %s", c.Int("t"), c.String("u"),
			)
			err = wd.WaitWithTimeout(
				page.Exists(c.String("s")), time.Duration(c.Int("t"))*time.Second,
			)
			if err != nil {
				return err
			}

			log.Printf("retrieving element: %s", c.String("s"))
			we, err := wd.FindElement(selenium.ByCSSSelector, c.String("s"))
			if err != nil {
				return err
			}

			log.Println("taking screenshot of element")
			b, err := we.Screenshot(false)
			if err != nil {
				return err
			}

			id := uuid.New()
			fp := fmt.Sprintf("/Users/revett/Downloads/%s.jpg", id.String())

			log.Printf("writing image to file: %s", fp)
			err = ioutil.WriteFile(fp, b, 0600)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
