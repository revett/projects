package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tebeka/selenium"
)

const (
	defaultTimeout              = 5
	defaultWebDriverBrowserName = "firefox"
)

var (
	browserName string
	host        string
	timeout     int

	rootCmd = &cobra.Command{
		Use:   "mapbox-screenshot",
		Short: "Take screenshots of Mapbox maps using Selenium",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&browserName,
		"browserName",
		defaultWebDriverBrowserName,
		"webdriver browsername",
	)
	rootCmd.PersistentFlags().StringVar(
		&host,
		"host",
		selenium.DefaultURLPrefix,
		"local selenium host uri",
	)
	rootCmd.PersistentFlags().IntVar(
		&timeout,
		"timeout",
		defaultTimeout,
		"timeout (seconds) to wait for page/element to load",
	)
}
