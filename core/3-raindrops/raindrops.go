// Package raindrops solves the Raindrops exercise at exercism.io
package raindrops

import "fmt"

// Convert converts the given integer to it's corresponding string
func Convert(in int) string {
	var s string

	if in%3 == 0 {
		s += "Pling"
	}

	if in%5 == 0 {
		s += "Plang"
	}

	if in%7 == 0 {
		s += "Plong"
	}

	// i.e., none of the above were true
	if s == "" {
		s = fmt.Sprint(in)
	}

	return s
}
