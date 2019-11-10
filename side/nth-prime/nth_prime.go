package prime

// isPrime returns true if the input is a Prime number.
func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// primeGen retuns a closure that generates the next Prime Number.
func primeGen() func() int {
	p := 1
	return func() int {
		for {
			p++
			if isPrime(p) {
				return p
			}
		}
	}
}

// Nth returns the Nth prime. ok is false if n is invalid.
func Nth(n int) (prime int, ok bool) {
	if n <= 0 {
		return
	}
	nextPrime := primeGen()
	for i := 0; i < n; i++ {
		prime = nextPrime()
	}
	return prime, true
}
