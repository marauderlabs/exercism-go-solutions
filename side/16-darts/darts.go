package darts

import "math"

// dist calculates the distance of x,y from the origin(0,0)
func dist(x, y float64) float64 {
	x2 := math.Pow(float64(x), 2)
	y2 := math.Pow(float64(y), 2)
	return math.Sqrt(x2 + y2)
}

// radius of the circles
const (
	outer  = 10
	middle = 5
	inner  = 1
)

// Score returns the dart board score for the given x,y co-ord of the dart on the board
func Score(x, y float64) int {
	d := dist(x, y)

	switch {
	case d <= outer && d > middle:
		return 1
	case d <= middle && d > inner:
		return 5
	case d <= inner:
		return 10
	}

	return 0
}
