package utils

import (
	"math/rand"
	"strings"
)

var (
	Alphabet []rune
	AlphabetLen int
)

func SetupSlugs() {
	sb := strings.Builder{}
	for i := range(26) {
		sb.WriteRune(rune('a') + rune(i))
		sb.WriteRune(rune('A') + rune(i))
		if i <= 9 {
			sb.WriteRune(rune('0') + rune(i))
		}
	}
	sb.WriteString("-_")
	Alphabet = []rune(sb.String())
	AlphabetLen = len(Alphabet)
}

func GetSlug(l int) string {
	s := make([]rune, l)
	for i := range l {
		s[i] = Alphabet[rand.Intn(AlphabetLen)]
	}
	return string(s)
}
