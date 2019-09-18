package lsproduct

import (
	"errors"
	"unicode"
)

// productOfDigits returns the product of digits in s; error if invalid digit if found.
func productOfDigits(s string) (int64, error) {
	var prod int64 = 1
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return 0, errors.New("Invalid digit")
		}
		prod *= int64(r - '0')
	}
	return prod, nil
}

// LargestSeriesProduct returns the largest product for a contiguous substring of digits of given length.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	l := len(digits)
	if span > l || span < 0 {
		return 0, errors.New("Invalid length of span")
	}

	var max, prod int64
	var err error
	for i := 0; i <= l-span; i++ {

		sub := digits[i : i+span]
		if prod, err = productOfDigits(sub); err != nil {
			return 0, err
		}

		if prod > max {
			max = prod
		}
	}
	return max, nil
}
