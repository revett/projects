package language_test

import (
	"testing"

	"github.com/revett/projects/internal/pkg/language"
	"github.com/stretchr/testify/assert"
)

func TestWordConsonantVowelPattern(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"Closed": {
			input: "cat",
			want:  "CVC",
		},
	}

	for n, test := range tests {
		t.Run(n, func(subtest *testing.T) {
			w := language.Word("foo")
			assert.Equal(t, test.want, w.ConsonantVowelPattern())
		})
	}
}
