// Package proverb solves the proverb exercise at exercism.io
package proverb

import "fmt"

const pvLine string = "For want of a %v the %v was lost."
const pvLast string = "And all for the want of a %v."

// Proverb returns the proverb generated for the words given in rhyme
func Proverb(rhyme []string) []string {
	var pv []string

	l := len(rhyme)
	if l == 0 {
		return pv
	}

	for i := 0; i < l-1; i++ {
		pv = append(pv, fmt.Sprintf(pvLine, rhyme[i], rhyme[i+1]))
	}
	pv = append(pv, fmt.Sprintf(pvLast, rhyme[0]))
	return pv
}
