package page

import (
	"log"
	"time"

	"github.com/tebeka/selenium"
)

// Object is the page object model (POM) for a given web page.
type Object struct {
	wd selenium.WebDriver
}

// New creates a new Object struct with the provided WebDriver interface.
func New(wd selenium.WebDriver) Object {
	return Object{
		wd: wd,
	}
}

// FindElement returns a web element from the page, found using the provided
// CSS selector.
func (o Object) FindElement(s string) (selenium.WebElement, error) {
	log.Printf("retrieving element: %s", s)
	return o.wd.FindElement(selenium.ByCSSSelector, s)
}

// ScreenshotElement takes a screenshot of a web element on the page. The web
// element is found using a provided CSS selector.
func (o Object) ScreenshotElement(s string) ([]byte, error) {
	we, err := o.FindElement(s)
	if err != nil {
		return nil, err
	}

	log.Println("taking screenshot of element")

	return we.Screenshot(true)
}

// Visit instructs the browser to navigate to a provided URL.
func (o Object) Visit(u string) error {
	log.Printf("opening webpage: %s", u)
	return o.wd.Get(u)
}

// WaitForElement checks the web page for if a specific web element exists, at
// a regular internal. The web element is found using a provided CSS selector.
func (o Object) WaitForElement(s string, t int) error {
	log.Printf("waiting for element (up to %ds): %s", t, s)

	return o.wd.WaitWithTimeout(
		ElementExistsCondition(s), time.Duration(t)*time.Second,
	)
}
