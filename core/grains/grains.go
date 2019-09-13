package grains

import "errors"

// Square returns the number of grains on the square s.
func Square(s int) (uint64, error) {
	if s <= 0 || s > 64 {
		return 0, errors.New("Invalid square")
	}
	// simple multiplication (i.e., power of 2)
	return 1 << (uint(s) - 1), nil
}

// Total retuns the total number of grains on the whole board.
func Total() uint64 {
	return (1 << 64) - 1
}
