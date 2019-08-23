// Package collatzconjecture solves the 'Collatz Conjecture' exercise at exercism.io
package collatzconjecture

import "errors"

// CollatzConjecture returns the numbers of steps needed to reach 1
// by following Collatz Conjecture
func CollatzConjecture(in int) (int, error) {
	if in <= 0 {
		return 0, errors.New("Invalid input")
	}

	steps := 0
	for in != 1 {

		if in%2 == 0 {
			in /= 2
		} else {
			in = 3*in + 1
		}
		steps++
	}
	return steps, nil
}
