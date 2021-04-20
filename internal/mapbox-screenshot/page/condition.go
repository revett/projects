package page

import "github.com/tebeka/selenium"

// ElementExistsCondition returns a selenium.Condition which checks if a web
// element exists using a CSS selector.
func ElementExistsCondition(s string) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		om := New(wd)
		if _, err := om.FindElement(s); err != nil {
			return false, err
		}

		return true, nil
	}
}
