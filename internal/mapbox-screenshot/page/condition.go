package page

import (
	"fmt"

	"github.com/tebeka/selenium"
)

const localStorageKey = "mapbox.StylesLoaded"

// ElementExistsCondition returns a selenium.Condition which checks if a web
// element exists using a CSS selector.
func ElementExistsCondition(s string) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		s := fmt.Sprintf("return localStorage.getItem(\"%s\")", localStorageKey)

		v, err := wd.ExecuteScript(s, nil)
		if err != nil {
			return false, err
		}

		if v == "true" {
			return true, nil
		}

		return false, nil
	}
}
