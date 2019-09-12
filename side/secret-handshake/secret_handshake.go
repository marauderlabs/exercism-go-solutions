package secret

// a slice of functions to perform the required conversion
// if the n'th bit is set. Index on the slice corresponds to the bit.
var convert = []func([]string) []string{
	func(s []string) []string {
		return append(s, "wink")
	},
	func(s []string) []string {
		return append(s, "double blink")
	},
	func(s []string) []string {
		return append(s, "close your eyes")
	},
	func(s []string) []string {
		return append(s, "jump")
	},
	func(s []string) []string {
		// this reverses the slice in place but
		// we return it anyway to keep with the signature of `convert`
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s
	},
}

// Handshake returns the slice of the coded strings for the given binary
func Handshake(code uint) []string {

	var s []string
	var b uint
	// As long as some bit is set (>0)
	// check that bit, call the convert function to
	// convert it.
	// Then, shift right to get the next bit.
	// Also check only the least significant 5 bits (b<5).

	for code > 0 && b < 5 {
		if code&1 != 0 {
			s = convert[b](s)
		}
		code >>= 1
		b++
	}

	return s
}
