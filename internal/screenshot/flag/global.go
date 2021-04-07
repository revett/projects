package flag

import (
	"strconv"

	"github.com/tebeka/selenium"
	"github.com/urfave/cli/v2"
)

const (
	defaultTimeout              = 5
	defaultWebDriverBrowserName = "firefox"
	defaultWebDriverHost        = selenium.DefaultURLPrefix
)

// Global generates all of the global CLI flags.
func Global() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "browserName",
			Aliases:     []string{"bn"},
			DefaultText: defaultWebDriverBrowserName,
			Usage:       "webdriver browsername",
			Value:       defaultWebDriverBrowserName,
		},
		&cli.StringFlag{
			Name:        "host",
			Aliases:     []string{"ho"},
			DefaultText: defaultWebDriverHost,
			Usage:       "local webdriver host uri",
			Value:       defaultWebDriverHost,
		},
		&cli.IntFlag{
			Name:        "timeout",
			Aliases:     []string{"t"},
			DefaultText: strconv.Itoa(defaultTimeout),
			Usage:       "timeout (seconds) to wait for page/element to load",
			Value:       defaultTimeout,
		},
		&cli.StringFlag{
			Name:     "url",
			Aliases:  []string{"u"},
			Required: true,
			Usage:    "website to visit",
		},
	}
}
