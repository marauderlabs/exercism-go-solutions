// Package strain solves the strain problem at exercism.io
package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep returns the elements of `in` for which the predicate f returns true
func (in Ints) Keep(f func(int) bool) Ints {
	var out Ints
	for _, elt := range in {
		if f(elt) {
			out = append(out, elt)
		}
	}
	return out
}

// Discard returns the elements of `in` for which the predicate f returns false
func (in Ints) Discard(f func(int) bool) Ints {
	var out Ints
	for _, elt := range in {
		if !f(elt) {
			out = append(out, elt)
		}
	}
	return out
}

// Keep returns the elements of `in` for which the predicate f returns true
func (in Lists) Keep(f func([]int) bool) Lists {
	var out Lists
	for _, elt := range in {
		if f(elt) {
			out = append(out, elt)
		}
	}
	return out
}

// Keep returns the elements of `in` for which the predicate f returns true
func (in Strings) Keep(f func(string) bool) Strings {
	var out Strings
	for _, elt := range in {
		if f(elt) {
			out = append(out, elt)
		}
	}
	return out
}
