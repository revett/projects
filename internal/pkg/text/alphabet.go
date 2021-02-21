package text

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// EnglishAlphabet contains the 26 lowercase letters of the english alphabet.
const EnglishAlphabet = "abcdefghijklmnopqrstuvwxyz"

// Alphabet contains a slice of letters.
type Alphabet struct {
	Combinations []string
	Letters      []string
}

// NewAlphabet creates and configures an Alphabet struct.
func NewAlphabet(s string, opts ...func(*Alphabet)) *Alphabet {
	a := &Alphabet{
		Letters: strings.Split(s, ""),
	}

	for _, o := range opts {
		o(a)
	}

	a.generateCombinations()
	return a
}

// WithMixedCase duplicates the original alphabet string so that each letter is
// both upper and lower case.
func WithMixedCase() func(*Alphabet) {
	return func(a *Alphabet) {
		var mc []string
		for _, l := range a.Letters {
			mc = append(mc, strings.ToLower(l), strings.ToUpper(l))
		}

		a.Letters = mc
	}
}

// RandomLetterPair returns a random letter pair from the configured alphabet.
func (a Alphabet) RandomLetterPair() string {
	return a.Combinations[rand.Intn(len(a.Combinations))]
}

func (a *Alphabet) generateCombinations() {
	var c []string

	for _, lx := range a.Letters {
		for _, ly := range a.Letters {
			c = append(c, fmt.Sprintf("%s%s", lx, ly))
		}
	}

	a.Combinations = c
}
