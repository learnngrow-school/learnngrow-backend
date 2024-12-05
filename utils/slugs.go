package utils

import (
	"math/rand"
	"strings"
)

var (
	Alphabet []rune
	AlphabetLen int
)

func SetupSlugs() []rune {
	sb := strings.Builder{}
	for i := range(26) {
		sb.WriteRune(rune('a') + rune(i))
		sb.WriteRune(rune('A') + rune(i))
		if i <= 9 {
			sb.WriteRune(rune('0') + rune(i))
		}
	}
	sb.WriteString("-_")
	s := sb.String()
	AlphabetLen = len(s)
	return []rune(s)
}

func GetSlug(l int) string {
	s := make([]rune, l)
	for i := range l {
		s[i] = Alphabet[rand.Intn(AlphabetLen)]
	}
	return string(s)
}
