package fastfact

import (
	"errors"
	"math/big"
)

// miniFact returns factorial of numbers 0, 1, 2.
func miniFact(n uint64) (*big.Int, error) {
	// 0
	if n == 0 {
		return big.NewInt(1), nil
	}

	// 1 or 2
	if n == 1 || n == 2 {
		return big.NewInt(int64(n)), nil
	}

	// n > 2
	var retval *big.Int
	return retval, errors.New("n greater than 2")
}

// numToRanges uniformly splits given number to the
// number of segments specified by workerNum.
//
// This is used to split the work for concurrent calculations.
func numToRanges(num, workerNum uint64) [][]uint64 {
	ranges := make([][]uint64, workerNum)
	var start uint64 = 1
	var stop uint64
	var step uint64
	var incLastStop bool // increment flag

	switch {
	case num%workerNum == 0:
		step = num / workerNum
	case num%workerNum == 1:
		step = num / workerNum
		incLastStop = true
	default: // num % workerNum > 1
		step = num/workerNum + 1
	}

	for i := 0; uint64(i) < workerNum; i++ {
		rng := make([]uint64, 2)
		lastIteration := uint64(i) == workerNum-1
		stop = start + step - 1

		if lastIteration && incLastStop {
			stop++
		}
		if stop > num {
			stop = num
		}
		rng[0], rng[1] = start, stop
		start = stop + 1

		ranges[i] = rng
	}
	return ranges
}

// mulRangeIter returns the product of all
// integers in the range [a, b] inclusively.
// If a > b (empty range), the result is 1.
//
// Utilizes iterative multiplication, slow.
func mulRangeIter(a, b uint64) *big.Int {
	res := big.NewInt(1)
	if a > b {
		return res
	}
	for i := a; i <= b; i++ {
		res.Mul(res, big.NewInt(int64(i)))
	}
	return res
}

// mulRangeFast returns the product of all
// integers in the range [a, b] inclusively.
// If a > b (empty range), the result is 1.
//
// Utilizes big.Int.MulRange, fast.
func mulRangeFast(a, b uint64) *big.Int {
	res := big.NewInt(1)
	return res.MulRange(int64(a), int64(b))
}
