package atbash

import (
	"strings"
	"unicode"
)

// Atbash returns the atbash encoded string
func Atbash(s string) string {
	var b strings.Builder
	var l uint

	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			continue
		}

		if l != 0 && l%5 == 0 {
			b.WriteRune(' ')
		}
		l++

		if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
			b.WriteRune('z' - (r - 'a'))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}
