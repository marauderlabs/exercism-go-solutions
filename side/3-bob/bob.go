// Package bob solves the Bob exercise at exercism.io
package bob

import (
	"strings"
	"unicode"
)

type remarkType int

// Different remarks that Bob may be made on Bob
const (
	question remarkType = iota
	yell
	yellQuestion
	nothing
	anything
	maxRemark // invalid. unused for now
)

// initResponses initializes all possible responses that Bob provides for different remarks
func initResponses() []string {
	responses := make([]string, maxRemark)
	responses[question] = "Sure."
	responses[yell] = "Whoa, chill out!"
	responses[yellQuestion] = "Calm down, I know what I'm doing!"
	responses[nothing] = "Fine. Be that way!"
	responses[anything] = "Whatever."
	return responses
}

/*
func upper(remark string) bool {
	atLeastOneUpper := false
	for _, c := range remark {
		if c >= 'a' && c <= 'z' {
			return false
		}

		if c >= 'A' && c <= 'Z' {
			atLeastOneUpper = true
		}
	}
	return atLeastOneUpper
}
*/

func upper(remark string) bool {
	atLeastOneUpper := false
	for _, r := range remark {
		if unicode.IsLower(r) {
			return false
		}
		if unicode.IsUpper(r) {
			atLeastOneUpper = true
		}
	}
	return atLeastOneUpper
}

func getRemarkType(remark string) remarkType {
	remark = strings.TrimSpace(remark)

	// no input or just spaces
	if len(remark) == 0 {
		return nothing
	}

	// is question?
	q := remark[len(remark)-1] == '?'

	// is a yell? all upper case?
	y := upper(remark)

	switch {

	case q && y:
		return yellQuestion

	case q:
		return question

	case y:
		return yell
	}

	return anything
}

// Hey returns bob's responses to the remarks he gets
func Hey(remark string) string {
	responses := initResponses()
	remarkType := getRemarkType(remark)
	return responses[remarkType]
}
