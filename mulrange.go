package fastfact

import (
	"math/big"
)

// MulRangeFact calculates factorial of
// the given number by utilizing Int.MulRange
// for multiplication.
func MulRangeFact(n uint64) *big.Int {
	if n < 3 {
		retval, _ := miniFact(n)
		return retval
	}

	// n > 2
	res := big.NewInt(1)
	return res.MulRange(2, int64(n))
}
