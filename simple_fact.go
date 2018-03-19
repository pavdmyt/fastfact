package fastfact

import "math/big"

// SimpleFactIter calculates factorial of
// the given number by iterative multiplication.
//
// Utilizes slow backend for multiplication: mulRangeIter.
func SimpleFactIter(n uint64) *big.Int {
	if n < 3 {
		retval, _ := miniFact(n)
		return retval
	}
	// n > 2
	return mulRangeIter(1, n)
}

// SimpleFactFast calculates factorial of the
// given number by utilizing Int.MulRange for
// multiplication.
//
// Utilizes fast backend for multiplication: mulRangeFast.
func SimpleFactFast(n uint64) *big.Int {
	if n < 3 {
		retval, _ := miniFact(n)
		return retval
	}
	// n > 2
	return mulRangeFast(1, n)
}
