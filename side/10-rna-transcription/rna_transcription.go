// Package strand solves the RNA Transcription exercise at exercism.io
package strand

import "strings"

// ToRNA returns the RNA transcript of the given DNA
func ToRNA(dna string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 'G':
			return 'C'
		case 'C':
			return 'G'
		case 'T':
			return 'A'
		case 'A':
			return 'U'
		}
		return r
	}, dna)
}

// // ToRNA returns the RNA transcript of the given DNA
// // An alternate implementation with strings pkg
// func ToRNA(dna string) string {
// 	var rna string
// 	var r rune
// 	for _, d := range dna {
// 		switch d {
// 		case 'G':
// 			r = 'C'
// 		case 'C':
// 			r = 'G'
// 		case 'T':
// 			r = 'A'
// 		case 'A':
// 			r = 'U'
// 		default:
// 			r = d
// 		}
// 		rna += string(r)
// 	}

// 	return rna
// }
