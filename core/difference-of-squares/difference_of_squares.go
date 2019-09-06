// Package diffsquares solves the 'Difference Of Squares' exercise at exercism.io.
package diffsquares

// SquareOfSum returns the square of the sum of first n natural numbers.
func SquareOfSum(n int) int {
	s := (n * (n + 1)) / 2
	return s * s
}

// SumOfSquares returns the sum of the square of first n natural numbers.
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference returns the difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
