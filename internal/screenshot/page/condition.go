package page

import "github.com/tebeka/selenium"

// Exists returns a selenium.Condition which checks if a web element exists
// using a CSS selector.
func Exists(s string) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		if _, err := wd.FindElement(selenium.ByCSSSelector, s); err != nil {
			return false, err
		}

		return true, nil
	}
}
