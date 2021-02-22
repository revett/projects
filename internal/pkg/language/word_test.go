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
		"ClosedShort": {
			input: "at",
			want:  "VC",
		},
		"Open": {
			input: "me",
			want:  "CV",
		},
		"VowelConsonantE": {
			input: "made",
			want:  "CVCE",
		},
		"FalseVowelConsonantE": {
			input: "abductee",
			want:  "VCCVCCVV",
		},
		"RControlled": {
			input: "her",
			want:  "CVR",
		},
		"RControlledNested": {
			input: "bird",
			want:  "CVRC",
		},
		"Diphthong": {
			input: "meet",
			want:  "CVVC",
		},
		"ConsonantLE": {
			input: "table",
			want:  "CVCLE",
		},
		"FalseConsonantLE": {
			input: "tale",
			want:  "CVCC",
		},
		"UpperCase": {
			input: "FOOD",
			want:  "CVVC",
		},
		"MixedCase": {
			input: "Golang",
			want:  "CVCVCC",
		},
	}

	for n, test := range tests {
		t.Run(n, func(subtest *testing.T) {
			w := language.NewWord(test.input)
			assert.Equal(t, test.want, w.ConsonantVowelPattern())
		})
	}
}
