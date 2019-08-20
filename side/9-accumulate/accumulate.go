// Package accumulate solves the accumulate challenge at exercism.io
package accumulate

type converter func(string) string

// Accumulate is like the reduce method in many languages. It returns the
// accumulated result of the elements in the slice converted using the converter
func Accumulate(in []string, f converter) []string {
	var out []string
	for _, s := range in {
		out = append(out, f(s))
	}
	return out
}
