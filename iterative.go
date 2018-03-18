package fastfact

import (
	"math/big"
)

// IterFact calculates factorial of
// the given number by iterative
// multiplication.
func IterFact(n uint64) *big.Int {
	if n < 3 {
		retval, _ := miniFact(n)
		return retval
	}
	return mulRange(1, n)
}
