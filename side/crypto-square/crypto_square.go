package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// getRectSize returns the minimum size of the rectangle needed for the encryption.
func getRectSize(l int) (r, c int) {
	sqrt := int(math.Ceil(math.Sqrt(float64(l))))

	if l == sqrt*sqrt { // perfect square
		return sqrt, sqrt
	}

	if (sqrt-1)*sqrt >= l {
		// square of the root is quite bigger than needed
		return (sqrt - 1), sqrt
	}

	return sqrt, sqrt
}

// compact returns the compacted string by removing spaces
func compact(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

// Encode returns the crypto text for the given plain text.
func Encode(pt string) string {
	pt = compact(pt)
	l := len(pt)
	r, c := getRectSize(l)

	rect := make([]string, r, r)
	cur := 0

	// slice the compacted string and put them in a 2D array(slice of strings)
	for i := range rect {
		if cur+c > l {
			// the last snip can be padded if needed
			rect[i] = pt[cur:] + strings.Repeat(" ", l-(cur-c))
		} else {
			rect[i] = pt[cur : cur+c]
		}
		cur += c
	}

	var b strings.Builder
	// now walk it top down and build the cipher
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			b.WriteByte(rect[j][i])
		}
		if i != c-1 { // space between the words
			b.WriteRune(' ')
		}
	}
	return b.String()
}
