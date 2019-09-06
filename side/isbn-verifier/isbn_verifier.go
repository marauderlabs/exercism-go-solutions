package isbn

import "unicode"

// IsValidISBN returns true if the input is a valid ISBN 10 number; false otherwise.
func IsValidISBN(n string) bool {
	mul := 10 // the multiplier
	val := 0  // the digit value
	sum := 0

	for _, r := range n {

		if r == '-' { // ignore
			continue
		}

		if !unicode.IsDigit(r) {

			if r != 'X' { // any other character
				return false
			}
			if mul != 1 { // 'X' is not at the last position, as check digit
				return false
			}
			val = 10 // we have a check digit

		} else {
			val = int(r - '0') // what's the digit?
		}

		sum += (val * mul)
		mul--
	}

	// can't find 10 digits
	if mul != 0 {
		return false
	}

	return sum%11 == 0
}
