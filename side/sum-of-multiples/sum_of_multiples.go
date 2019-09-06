package summultiples

// SumMultiples retuns the sum of all the unique multiples of particular numbers up to but not including that number.
func SumMultiples(limit int, divisors ...int) int {
	sum := 0

	for i := 1; i < limit; i++ {
		for _, d := range divisors {

			if d == 0 { // don't divide by 0; also 0 is inconsequential
				continue
			}
			if i%d == 0 {
				sum += i
				break
			}
		}
	}

	return sum
}
