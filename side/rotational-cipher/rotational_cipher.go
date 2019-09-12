package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher returns the plaintext p encrypted by rotating r times
func RotationalCipher(p string, r int) string {
	// just a shortcut!
	if r == 0 || r == 26 {
		return p
	}

	return strings.Map(func(c rune) rune {
		k := rune(r)

		switch {
		case unicode.IsLower(c):
			return (c-'a'+k)%26 + 'a'

		case unicode.IsUpper(c):
			return (c-'A'+k)%26 + 'A'

		default: // untouched
			return c
		}
	}, p)
}
