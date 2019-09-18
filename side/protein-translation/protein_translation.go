// Package protein solves the side exercise 'Protein Translation' at exercism.io
package protein

import (
	"errors"
	"strings"
)

var codonMap map[string]string

// just for easy copy-paste initialization from the README
type c2p struct {
	codons  string
	protein string
}

const stopProtein = "STOP"

var codon2protein = []c2p{
	{"AUG", "Methionine"},
	{"UUU, UUC", "Phenylalanine"},
	{"UUA, UUG", "Leucine"},
	{"UCU, UCC, UCA, UCG", "Serine"},
	{"UAU, UAC", "Tyrosine"},
	{"UGU, UGC", "Cysteine"},
	{"UGG", "Tryptophan"},
	{"UAA, UAG, UGA", stopProtein},
}

// tokenize the get the codons and map them to the protein given
func initCodonMap() map[string]string {
	m := make(map[string]string)
	for _, e := range codon2protein {
		for _, c := range strings.Split(e.codons, ", ") {
			m[c] = e.protein
		}
	}
	return m
}

// Some errors expected from us

// ErrStop means only the STOP codon is present in the sequence
var ErrStop = errors.New("Only STOP codon present")

// ErrInvalidBase indicated the presence of an invalid base codon in the sequence
var ErrInvalidBase = errors.New("Invalid base codon")

// FromCodon returns the protein name for the given codon
func FromCodon(codon string) (string, error) {
	// init the map once
	if codonMap == nil {
		codonMap = initCodonMap()
	}

	var p string
	var found bool

	if p, found = codonMap[codon]; !found {
		return "", ErrInvalidBase
	}

	if p == stopProtein {
		return "", ErrStop
	}

	return p, nil
}

// FromRNA returns the Protein for the given RNA
func FromRNA(r string) ([]string, error) {
	l := len(r)

	var p []string // the protein
	for i := 0; i < l; i += 3 {
		codon := r[i : i+3]

		a, err := FromCodon(codon)

		if err == ErrStop {
			break
		}

		if err != nil {
			return p, err
		}

		p = append(p, a)
	}
	return p, nil
}
