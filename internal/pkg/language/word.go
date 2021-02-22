package language

import (
	"strings"
)

// Word holds linguistic functions for analysis.
type Word struct {
	Letters []string
	Text    string
}

// NewWord creates a new word.
func NewWord(s string) *Word {
	t := strings.ToLower(s)
	return &Word{
		Letters: strings.Split(t, ""),
		Text:    t,
	}
}

// ConsonantVowelPattern translates the word in to a CVC syllable pattern.
func (w Word) ConsonantVowelPattern() string {
	var p []string

	v := vowels()
	for _, l := range w.Letters {
		if _, ok := v[l]; ok {
			p = append(p, "V")
			continue
		}

		p = append(p, "C")
	}

	l := len(p)

	if strings.Contains(w.Text, "r") {
		for i, e := range p {
			if i+1 == l {
				break
			}

			if e == "V" && w.Letters[i+1] == "r" {
				p[i+1] = "R"
			}
		}
	}

	if len(w.Text) < 3 {
		return strings.Join(p, "")
	}

	if strings.HasSuffix(w.Text, "le") && p[l-3] == "C" {
		p[l-2] = "L"
		p[l-1] = "E"
		return strings.Join(p, "")
	}

	if strings.HasSuffix(w.Text, "e") && p[l-3] == "V" && p[l-2] == "C" {
		p[l-1] = "E"
	}

	return strings.Join(p, "")
}

func vowels() map[string]int {
	return map[string]int{
		"a": 0,
		"e": 0,
		"i": 0,
		"o": 0,
		"u": 0,
	}
}
