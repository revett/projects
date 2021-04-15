package browser_test

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/revett/projects/internal/screenshot/browser"
	"github.com/stretchr/testify/require"
	"github.com/tebeka/selenium"
)

func TestNew(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		m   mock
		err bool
	}{
		"Success": {},
		"Error": {
			m: mock{
				err: errors.New("error"),
			},
			err: true,
		},
	}

	for n, tc := range tests {
		tc := tc

		t.Run(n, func(t *testing.T) {
			t.Parallel()
			_, err := browser.New(tc.m.Remoter, "firefox", "http://example.com")
			require.Equal(t, tc.err, err != nil)
		})
	}
}

type mock struct {
	err error
}

func (m mock) Remoter(c selenium.Capabilities, u string) (selenium.WebDriver, error) {
	return nil, m.err
}

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}
