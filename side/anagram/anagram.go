package anagram

import (
	"strings"
	"unicode"
)

const alphabets = 26

type letterCount [alphabets]int

// countLetters returns an array of integers containing the count of letters in given string.
// count of 'a' is index 0 and 'z' at 25
func countLetters(s string) letterCount {
	var count letterCount
	for _, r := range s {
		if unicode.IsLetter(r) {
			count[r-'a']++
		}
	}
	return count
}

// isEqual returns true if the counts at each index in c1 and c2 match
func isEqual(c1, c2 letterCount) bool {
	for i := 0; i < alphabets; i++ {
		if c1[i] != c2[i] {
			return false
		}
	}
	return true
}

// Detect returns a slice of strings from candidates that are an anagram of sub
func Detect(sub string, candidates []string) []string {
	var r []string
	sub = strings.ToLower(sub)
	sc := countLetters(sub)

	for i, c := range candidates {
		c = strings.ToLower(c)

		if sub == c { // impostor
			continue
		}

		cc := countLetters(c)

		if isEqual(sc, cc) {
			r = append(r, candidates[i]) // append the original, as c is now lowercase
		}
	}

	return r
}
