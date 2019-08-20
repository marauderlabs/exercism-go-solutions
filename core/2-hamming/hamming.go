// Package hamming solves the hamming exercise at exercism.io
package hamming

import "errors"

// Distance returns the hamming Distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strings are unequal")
	}

	la := len(a)
	dist := 0
	for i := 0; i < la; i++ {
		if a[i] != b[i] {
			dist++
		}
	}

	return dist, nil
}
