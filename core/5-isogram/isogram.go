// Package isogram sovles the isogram exercise at exercism.io
package isogram

import "unicode"

// IsIsogram returns true if the string s is an Isogram; false otherwise
func IsIsogram(s string) bool {
	foundRune := [26]bool{} //'a' to 'z'

	for _, r := range s {

		if !unicode.IsLetter(r) {
			continue
		}

		// convert the rune to lowercase to index `foundRune` with a simple `- 'a'`.
		r = unicode.ToLower(r)
		i := r - 'a'

		if foundRune[i] == true {
			return false
		}
		foundRune[i] = true
	}

	return true
}
