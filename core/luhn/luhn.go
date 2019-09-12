package luhn

import (
	"errors"
	"strings"
	"unicode"
)

// clean returns the string, with spaces removed from it.
// Returns error if non-numerical runes other than spaces are found.
func clean(s string) (string, error) {
	var b strings.Builder

	for _, r := range s {
		if unicode.IsSpace(r) { // Don't include spaces
			continue
		}
		if !unicode.IsNumber(r) { // NaN
			return "", errors.New("invalid")
		}
		b.WriteRune(r)
	}

	return b.String(), nil
}

// Valid return true if the input is a Valid Luhn number.
func Valid(s string) bool {
	c, err := clean(s)

	if err != nil { // found something dirty
		return false
	}

	l := len(c) // too short
	if l <= 1 {
		return false
	}

	sum := 0

	for i := 0; i < l; i++ {
		// get the char from the right and make a digit out of it
		val := int(c[l-i-1] - '0')

		if i%2 == 1 { // every other second digit
			val *= 2
			if val > 9 {
				val -= 9
			}
		}

		sum += val
	}

	return sum%10 == 0
}
