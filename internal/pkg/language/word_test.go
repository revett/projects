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
			input: "le",
			want:  "CV",
		},
		"UpperCase": {
			input: "FOOD",
			want:  "CVVC",
		},
		"MixedCase": {
			input: "Golang",
			want:  "CVCVCC",
		},
		"RControlledConsonantLE": {
			input: "hurdle",
			want:  "CVRCLE",
		},
		"RControlledVowelConsonantE": {
			input: "hurade",
			want:  "CVRVCE",
		},
	}

	for n, tc := range tests {
		t.Run(n, func(t *testing.T) {
			w := language.NewWord(tc.input)
			assert.Equal(t, tc.want, w.ConsonantVowelPattern())
		})
	}
}
