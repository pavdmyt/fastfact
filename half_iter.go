package fastfact

import (
	"math/big"
)

// HalfIterFact calculates factorial of
// the given number by reducing the number
// of multiplications by factor of 2.
//
// Details:
// http://pavdmyt.com/digging-around-factorial-function/
func HalfIterFact(n uint64) *big.Int {
	// 0
	one := big.NewInt(1)
	if n == 0 {
		return one
	}

	// 1 or 2
	bigN := big.NewInt(int64(n))
	if n == 1 || n == 2 {
		return bigN
	}

	// other
	var res *big.Int

	// Handle odd input
	if n%2 == 1 {
		res = big.NewInt(int64(n))
		bigN.Sub(bigN, one) // bigN -= 1
	} else {
		res = one
	}

	zero := big.NewInt(0)
	two := big.NewInt(2)
	i := big.NewInt(0)

	for bigN.Cmp(zero) == 1 { // bigN > 0
		i.Add(i, bigN)      // i += bigN
		res.Mul(res, i)     // res *= i
		bigN.Sub(bigN, two) // bigN -= 2
	}
	return res
}
