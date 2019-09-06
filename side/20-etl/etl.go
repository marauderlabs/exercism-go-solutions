package etl

import "strings"

// Transform returns into the new map for the legacy map
func Transform(old map[int][]string) map[string]int {
	n := make(map[string]int)

	for score, letters := range old {
		for _, c := range letters {
			l := strings.ToLower(c) // it's just 1 char/rune each but the test uses string. So shall we.
			n[l] = score
		}
	}

	return n
}
