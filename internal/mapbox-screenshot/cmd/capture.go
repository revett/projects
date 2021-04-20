package cmd

import (
	"fmt"
	"log"
	"net/url"

	"github.com/revett/projects/internal/mapbox-screenshot/browser"
	"github.com/revett/projects/internal/mapbox-screenshot/imgio"
	"github.com/revett/projects/internal/mapbox-screenshot/page"
	"github.com/spf13/cobra"
	"github.com/tebeka/selenium"
)

const (
	urlArgPosition = iota
	selectorArgPosition

	requiredCaptureArgs = 2
)

var captureCmd = &cobra.Command{
	Use:   "capture [url] [selector]",
	Short: "Screenshot a map",
	Args:  validateArgs,
	RunE:  captureElements,
}

func init() {
	rootCmd.AddCommand(captureCmd)
}

func captureElements(cmd *cobra.Command, args []string) error {
	url := args[urlArgPosition]
	selector := args[selectorArgPosition]

	b, err := browser.New(selenium.NewRemote, browserName, host)
	if err != nil {
		return err
	}

	defer func() {
		if err := b.Quit(); err != nil {
			log.Println("err")
		}
	}()

	om := page.New(b)
	if err = om.Visit(url); err != nil {
		return err
	}

	err = om.WaitForElement(selector, timeout)
	if err != nil {
		return err
	}

	bytes, err := om.ScreenshotElement(selector)
	if err != nil {
		return err
	}

	return imgio.Write(bytes)
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) != requiredCaptureArgs {
		return fmt.Errorf("requires exactly %d arguments", requiredCaptureArgs)
	}

	if _, err := url.ParseRequestURI(args[urlArgPosition]); err != nil {
		return err
	}

	return nil
}
