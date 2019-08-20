// Package acronym solves the acronym exercise at exercism.io
package acronym

import "strings"

// Abbreviate returns the acronym of the given phrase.
func Abbreviate(s string) string {
	// tokenize the string based on the space and hypens
	tokens := strings.FieldsFunc(s, func(r rune) bool {
		return (r == ' ' || r == '-')
	})

	acr := ""
	// lets iterate over the tokens and use the first chars
	// to create acronym.
	for _, tok := range tokens {

		// now for each token, chomp off any underscores on either end
		tok = strings.TrimFunc(tok, func(r rune) bool {
			return r == '_'
		})

		// concat the first chars.
		acr += string(tok[0])
	}

	// promote them to uppercase
	return strings.ToUpper(acr)
}
