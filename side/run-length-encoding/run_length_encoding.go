package encode

import (
	"fmt"
	"strings"
	"unicode"
)

// RunLengthEncode returns the run-length encoded input
func RunLengthEncode(in string) string {
	var out strings.Builder

	l := len(in)

	if l == 0 {
		return ""
	}

	var p byte
	var c int

	for i := 0; i < l; i++ {
		if in[i] != p {
			if p != 0 {
				encode(&out, p, c)
			}
			c = 0
		}
		p = in[i]
		c++
	}

	encode(&out, p, c)

	return out.String()
}

func encode(out *strings.Builder, b byte, c int) {
	var s string
	if c == 1 {
		s = fmt.Sprintf("%s", string(b))
	} else {
		s = fmt.Sprintf("%v%s", c, string(b))
	}
	out.WriteString(s)
}

//RunLengthDecode returns the run-length decoded input
func RunLengthDecode(in string) string {
	var n int
	var out strings.Builder
	for _, r := range in {
		if unicode.IsNumber(r) {
			n = n*10 + int(r-'0')
		} else {
			if n > 0 {
				out.WriteString(strings.Repeat(string(r), n))
			} else {
				out.WriteRune(r)
			}
			n = 0
		}
	}
	return out.String()
}
