package perfect

import "errors"

// Classification is Nicomachus classification of a given number based on aliquot sum
type Classification int

const (
	// ClassificationDeficient means the number is less than the sum.
	// ClassificationPerfect means it's equal to the sum.
	// ClassificationAbundant  means the number is greater than the sum.
	ClassificationDeficient = iota
	ClassificationPerfect
	ClassificationAbundant
)

// ErrOnlyPositive is returned by Classify if inputs aren't positive.
var ErrOnlyPositive = errors.New("Only positive numbers are allowed")

// getFactors find the factors of the given number
func getFactors(n int64) []int64 {
	var f []int64
	var i int64

	for i = 1; i <= n/2; i++ {
		d := n / i
		if (n%i == 0) && (i*d == n) {
			f = append(f, i)
		}
	}
	return f
}

// sum the ints in the given slice. A simple reduce.
func intSliceSum(n []int64) int64 {
	var sum int64
	for _, e := range n {
		sum += e
	}
	return sum
}

// Classify returns the Nicomachus classification of n.
func Classify(n int64) (Classification, error) {

	if n <= 0 {
		return 0, ErrOnlyPositive
	}

	f := getFactors(n)
	s := intSliceSum(f)

	if n < s {
		return ClassificationAbundant, nil
	} else if n == s {
		return ClassificationPerfect, nil
	} else {
		return ClassificationDeficient, nil
	}
}
