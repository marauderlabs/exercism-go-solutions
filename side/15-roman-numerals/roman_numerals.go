// Package romannumerals solves the 'Roman Numerals' side exercise at exercism.io
package romannumerals

import (
	"errors"
	"fmt"
	"strings"
)

// ToRomanNumeral returns the roman numeral string for the arabic input
func ToRomanNumeral(n int) (string, error) {
	var s string

	if n <= 0 || n > 3000 {
		return "", errors.New("Invalid input")
	}

	// There won't be need for the five and ten char as Num < 3k.
	// Passing them as it's mandatory.
	s += toRoman(n/1000, "M", "", "")
	s += toRoman(n/100%10, "C", "D", "M")
	s += toRoman(n/10%10, "X", "L", "C")
	s += toRoman(n%10, "I", "V", "X")

	return s, nil
}

func toRoman(n int, one, five, ten string) string {
	switch n {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		return strings.Repeat(one, n)
	case 4:
		return fmt.Sprintf("%v%v", one, five)
	case 5:
		return fmt.Sprintf("%v", five)
	case 6:
		fallthrough
	case 7:
		fallthrough
	case 8:
		return fmt.Sprintf("%v%v", five, strings.Repeat(one, n-5))
	case 9:
		return fmt.Sprintf("%v%v", one, ten)
	}
	return "" /* case 0 */
}
