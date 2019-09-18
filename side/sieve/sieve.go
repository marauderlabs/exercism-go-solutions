package sieve

type entry struct {
	n int  // the number
	p bool // is prime
}

// Sieve returns a list of prime numbers below limit
func Sieve(limit int) []int {
	var list []entry
	var result []int

	// take a list of all the numbers
	for i := 2; i <= limit; i++ {
		list = append(list, entry{i, true})
	}

	for _, e := range list {
		// I'm not a prime
		if !e.p {
			continue
		}

		// I'm a prime. Mark off my multiples. Start from 2, as x1 is myself.
		result = append(result, e.n)
		for j := 2; j*e.n <= limit; j++ {
			list[(j*e.n)-2].p = false // the slice's first element is 2, at index 0. so we offset by 2.
		}
	}

	return result
}
