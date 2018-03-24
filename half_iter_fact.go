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
	if n < 3 {
		retval, _ := miniFact(n)
		return retval
	}

	// n > 2
	one := big.NewInt(1)
	bigN := big.NewInt(int64(n))
	var total *big.Int

	// Handle odd input
	if n%2 == 1 {
		total = big.NewInt(int64(n))
		bigN.Sub(bigN, one) // bigN -= 1
	} else {
		total = one
	}

	zero := big.NewInt(0)
	two := big.NewInt(2)
	i := big.NewInt(0)

	for bigN.Cmp(zero) == 1 { // bigN > 0
		i.Add(i, bigN)      // i += bigN
		total.Mul(total, i) // total *= i
		bigN.Sub(bigN, two) // bigN -= 2
	}
	return total
}
