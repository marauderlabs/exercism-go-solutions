// Package leap solves the leap-year problem at exercism.io
package leap

// IsLeapYear returns true for leap years
func IsLeapYear(year int) bool {

	if (year%100 == 0) && (year%400 == 0) {
		return true
	}

	if year%100 == 0 && year%4 == 0 {
		return false
	}

	if year%4 == 0 {
		return true
	}

	return false
}
