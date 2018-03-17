package fastfact

import (
	"math/big"
)

// IterFact calculates factorial of
// the given number by iterative
// multiplication.
func IterFact(n uint64) *big.Int {
	one := big.NewInt(1)
	if n == 0 {
		return one
	}

	bigN := big.NewInt(int64(n))
	if n == 1 || n == 2 {
		return bigN
	}

	res := big.NewInt(1)
	i := big.NewInt(2)

	for i.Cmp(bigN) != 1 { // i <= bigN
		res.Mul(res, i) // res *= i
		i.Add(i, one)   // i++
	}
	return res
}
