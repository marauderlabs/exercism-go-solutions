// Package scrabble solves the scrabble exercise at exercism.io
package scrabble

import (
	"strings"
)

const alphabets = 26

// initPoint inits the scores per alphabet
func initPoints() []int {
	return []int{
		1,  //A
		3,  //B
		3,  //C
		2,  //D
		1,  //E
		4,  //F
		2,  //G
		4,  //H
		1,  //I
		8,  //J
		5,  //K
		1,  //L
		3,  //M
		1,  //N
		1,  //O
		3,  //P
		10, //Q
		1,  //R
		1,  //S
		1,  //T
		1,  //U
		4,  //V
		4,  //W
		8,  //X
		4,  //Y
		10, //Z
	}
}

// Score returns the score for a given word
func Score(s string) int {

	points := initPoints()
	score := 0
	// make them lower case, so we can lookup our scores
	s = strings.ToLower(s)

	for _, c := range s {
		index := int(c - 'a')
		score += points[index]
	}

	return score
}
