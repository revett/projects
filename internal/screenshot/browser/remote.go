package browser

import (
	"log"

	"github.com/tebeka/selenium"
)

type Remoter func(selenium.Capabilities, string) (selenium.WebDriver, error)

// New returns a remote Selenium client, using the provided host.
func New(r Remoter, bn string, h string) (selenium.WebDriver, error) {
	log.Printf("setting webdriver browsername: %s", bn)

	caps := selenium.Capabilities{
		"browserName": bn,
	}

	log.Printf("connecting to local selenium host: %s", h)

	return r(caps, h)
}
