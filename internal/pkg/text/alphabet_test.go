package text_test

import (
	"testing"

	"github.com/revett/projects/internal/pkg/text"
	"github.com/stretchr/testify/assert"
)

func TestAlphabetLetters(t *testing.T) {
	tests := map[string]struct {
		input         string
		want          []string
		withMixedCase bool
	}{
		"WithoutMixedCase": {
			input:         "xyz",
			want:          []string{"x", "y", "z"},
			withMixedCase: false,
		},
		"WithMixedCase": {
			input:         "xyz",
			want:          []string{"x", "y", "z", "X", "Y", "Z"},
			withMixedCase: true,
		},
		"WithMixedCaseUpperCaseInput": {
			input:         "XYZ",
			want:          []string{"x", "y", "z", "X", "Y", "Z"},
			withMixedCase: true,
		},
	}

	for n, test := range tests {
		t.Run(n, func(subtest *testing.T) {
			var opts []func(*text.Alphabet)
			if test.withMixedCase {
				opts = append(opts, text.WithMixedCase())
			}

			a := text.NewAlphabet(test.input, opts...)
			assert.ElementsMatch(t, a.Letters, test.want)
		})
	}
}

func TestAlphabetCombinations(t *testing.T) {
	tests := map[string]struct {
		input         string
		want          []string
		withMixedCase bool
	}{
		"WithoutMixedCase": {
			input:         "ab",
			want:          []string{"aa", "ab", "ba", "bb"},
			withMixedCase: false,
		},
		"WithMixedCase": {
			input: "ab",
			want: []string{
				"aa", "aA", "Aa", "AA", "ab", "aB", "Ab", "AB",
				"ba", "bA", "Ba", "BA", "bb", "bB", "Bb", "BB",
			},
			withMixedCase: true,
		},
	}

	for n, test := range tests {
		t.Run(n, func(subtest *testing.T) {
			var opts []func(*text.Alphabet)
			if test.withMixedCase {
				opts = append(opts, text.WithMixedCase())
			}

			a := text.NewAlphabet(test.input, opts...)
			assert.ElementsMatch(t, a.Combinations, test.want)
		})
	}
}
