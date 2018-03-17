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

	// n > 2
	one := big.NewInt(1)
	bigN := big.NewInt(int64(n))
	res := big.NewInt(1)
	i := big.NewInt(2)

	for i.Cmp(bigN) != 1 { // i <= bigN
		res.Mul(res, i) // res *= i
		i.Add(i, one)   // i++
	}
	return res
}
