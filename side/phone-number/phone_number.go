package phonenumber

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func isValidChar(r rune) bool {
	return (unicode.IsDigit(r) ||
		r == ' ' ||
		r == '.' ||
		r == '(' ||
		r == ')' ||
		r == '-' ||
		r == '+')
}

// Number returns a cleaned up telephone number and error if any
func Number(s string) (string, error) {
	var b strings.Builder
	// cleanup
	for _, r := range s {
		if !isValidChar(r) {
			e := fmt.Sprint("Invalid character in number:", r)
			return "", errors.New(e)
		}
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	n := b.String()

	if len(n) == 11 {
		if n[0] != '1' {
			return "", errors.New("Country code is not 1")
		}
		n = n[1:] // chop off the country code
	}

	if len(n) != 10 {
		e := fmt.Sprint("Invalid length:", n)
		return "", errors.New(e)
	}

	if n[0] == '0' || n[0] == '1' {
		return "", errors.New("Invalid area code")
	}

	if n[3] == '0' || n[3] == '1' {
		return "", errors.New("Invalid exchange code")
	}

	return n, nil
}

// AreaCode returns the area code of the given phone number
func AreaCode(s string) (string, error) {
	n, err := Number(s)
	if err != nil {
		return n, err
	}
	return n[:3], nil
}

// Format returns the formatted phone number as (xxx) xxx-xxxx
func Format(s string) (string, error) {
	n, err := Number(s)
	if err != nil {
		return n, err
	}
	return fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:]), nil
}
