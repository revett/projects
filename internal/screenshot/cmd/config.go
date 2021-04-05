package cmd

// Config holds configuration values for the screenshot command.
type Config struct {
	ElementSelector      string
	Timeout              int
	WebDriverBrowserName string
	WebDriverHost        string
	URL                  string
}

// IsElementScreenshot checks if the command is configured to take a screenshot
// of a page or element.
func (c Config) IsElementScreenshot() bool {
	return c.ElementSelector != ""
}
