package fastfact

import (
	"math/big"
)

// IterFact calculates factorial of
// the given number by iterative
// multiplication.
func IterFact(n uint64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	bigN := big.NewInt(int64(n))
	if n == 1 || n == 2 {
		return bigN
	}

	res := big.NewInt(1)
	one := big.NewInt(1)
	i := big.NewInt(2)

	for i.Cmp(bigN) != 1 {
		res.Mul(res, i)
		i.Add(i, one)
	}
	return res
}
