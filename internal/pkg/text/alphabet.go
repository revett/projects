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
	combinations []string
	letters      []string
}

// NewAlphabet creates and configures an Alphabet struct.
func NewAlphabet(s string, opts ...func(*Alphabet)) *Alphabet {
	a := &Alphabet{
		letters: strings.Split(s, ""),
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
		for _, l := range a.letters {
			mc = append(mc, strings.ToLower(l), strings.ToUpper(l))
		}

		a.letters = mc
	}
}

// RandomLetterPair returns a random letter pair from the configured alphabet.
func (a Alphabet) RandomLetterPair() string {
	return a.combinations[rand.Intn(len(a.combinations))]
}

func (a *Alphabet) generateCombinations() {
	var c []string

	for _, lx := range a.letters {
		for _, ly := range a.letters {
			c = append(c, fmt.Sprintf("%s%s", lx, ly))
		}
	}

	a.combinations = c
}
