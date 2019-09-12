package cipher

import (
	"strings"
	"unicode"
)

// Cipher implements Encode and Decode methods for crypto
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type shiftCipherKey []rune

// NewCaesar retuns a new Cipher that encode in Caesar cipher (shift - 3)
func NewCaesar() Cipher {
	return shiftCipherKey{3}
}

// NewShift retuns a new Shift Chiper that encodes by the given shift
func NewShift(n int) Cipher {
	if n < -25 || n > 25 || n == 0 {
		return nil
	}
	// negative? Rotate it!
	if n < 0 {
		n = 26 + n
	}
	return shiftCipherKey{rune(n)}
}

// NewVigenere returns a new Shift Chiper that
// encodes by the given shift distances in the key.
// transform the runes into their corresponding shifts
func NewVigenere(key string) Cipher {
	lk := len(key)
	xkey := make(shiftCipherKey, lk, lk)
	onlyA := true
	for i, k := range key {
		if unicode.IsLower(k) {
			if k != 'a' { // only As in the key. Invalid!
				onlyA = false
			}
			xkey[i] = k - 'a'
		} else { // anything that's not lower is invalid too
			return nil
		}
	}
	if onlyA {
		return nil
	}
	return xkey
}

// shift a given rune r by they key k
func shift(r, k rune) rune {
	return (r+k-'a')%26 + 'a'
}

// encode and decode are similar except for the shift.
// they can be unified into one code for better maintainability. But this works fine for this piece
func encode(k shiftCipherKey, p string) string {
	var b strings.Builder
	i := 0
	kl := len(k)
	for _, r := range p {
		if !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToLower(r)
		key := k[i%kl]
		i++
		b.WriteRune(shift(r, key))
	}
	return b.String()
}

func decode(k shiftCipherKey, p string) string {
	var b strings.Builder
	i := 0
	kl := len(k)
	for _, r := range p {
		if !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToLower(r)
		key := 26 - k[i%kl] // decode is just encode with inverted key (reverse shift)
		i++
		b.WriteRune(shift(r, key))
	}
	return b.String()
}

func (k shiftCipherKey) Encode(p string) string {
	return encode(shiftCipherKey(k), p)
}

func (k shiftCipherKey) Decode(p string) string {
	return decode(shiftCipherKey(k), p)
}
