// Package triangle solves the triangle problem at exercism.io
package triangle

import "math"

// Kind is the type of the triangle
type Kind int

const (
	// These are the different `Kind` of triangles
	NaT Kind = iota // NaT is not a triangle
	Equ             // Equ is Equilateral
	Iso             // Iso is Isoscles
	Sca             // Scalene - Remaining type
)

// KindFromSides returns the type of the triangle based on the sides
func KindFromSides(a, b, c float64) Kind {

	switch {

	// negative side
	case (a <= 0) || (b <= 0) || (c <= 0):
		return NaT

	// Not a number
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		return NaT

	// infinity
	case math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		return NaT

	// degenerate
	case ((a + b) < c) || ((a + c) < b) || ((b + c) < a):
		return NaT

	// equilateral
	case a == b && b == c:
		return Equ

	// isosceles
	case a == b || b == c || a == c:
		return Iso

	// Scalene
	default:
		return Sca
	}
}
