package pangram

import "unicode"

const alphabetSize = 26

// IsPangram returns true if the string is a pangram
func IsPangram(s string) bool {
	var foundRune [alphabetSize]bool
	var unique int

	for _, r := range s {

		if !unicode.IsLetter(r) {
			continue
		}

		if 'A' <= r && r <= 'Z' { // upper to lower
			r += 'a' - 'A'
		}

		if !foundRune[r-'a'] {
			foundRune[r-'a'] = true
			unique++
		}

		if unique == alphabetSize {
			return true
		}
	}

	return false
}
