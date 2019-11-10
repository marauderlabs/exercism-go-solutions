// Package pythagorean solves the Pythogorean triplet problem.
// A naive brute-force approach.
package pythagorean

// Triplet stores a pythogorean triplet.
type Triplet [3]int

// Range returns a slice of pythogorean triplets within the range min <-> max.
func Range(min, max int) []Triplet {
	var t []Triplet
	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			for k := j; k <= max; k++ {
				x := i * i
				y := j * j
				z := k * k
				if x+y == z || x+z == y || y+z == x {
					t = append(t, Triplet{i, j, k})
				}
			}
		}
	}
	return t
}

// Sum returns a slice of pythogreaon triplets whose sum is p.
func Sum(p int) []Triplet {
	max := p

	var t []Triplet
	for i := 1; i <= max; i++ {
		for j := i; j <= max; j++ {
			for k := j; k <= max; k++ {
				x := i * i
				y := j * j
				z := k * k
				if (x+y == z || x+z == y || y+z == x) && (i+j+k == p) {
					t = append(t, Triplet{i, j, k})
				}
			}
		}
	}
	return t
}
