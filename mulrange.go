package fastfact

import (
	"math/big"
)

// MulRangeFact calculates factorial of
// the given number by utilizing Int.MulRange
// for multiplication.
func MulRangeFact(n uint64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	bigN := big.NewInt(int64(n))
	if n == 1 || n == 2 {
		return bigN
	}

	res := big.NewInt(1)
	return res.MulRange(2, int64(n))
}
