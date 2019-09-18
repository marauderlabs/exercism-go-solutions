package diffiehellman

import "math/big"
import cr "crypto/rand"

// PrivateKey returns a random number in range [2, p)
func PrivateKey(p *big.Int) *big.Int {
	// generate between 0 and p-2 and add 2 later
	h := big.NewInt(0)
	h = h.Sub(p, big.NewInt(2))
	r, _ := cr.Int(cr.Reader, h)
	return r.Add(r, big.NewInt(2))
}

// PublicKey retuns the public Key for the given primes p and G
// and the private key a
func PublicKey(a, p *big.Int, g int64) *big.Int {
	// A = g**a mod p
	bg := big.NewInt(g)
	A := big.NewInt(0)
	return A.Exp(bg, a, p)
}

// SecretKey retuns a secret key given
// the other's Public key B,
// self private key a and
// the prime number p
func SecretKey(a, B, p *big.Int) *big.Int {
	// s = B**a mod p
	s := big.NewInt(0)
	return s.Exp(B, a, p)
}

// NewPair returns a pair of Private and Public keys
func NewPair(p *big.Int, g int64) (k, K *big.Int) {
	k = PrivateKey(p)
	K = PublicKey(k, p, g)
	return
}
