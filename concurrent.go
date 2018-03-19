package fastfact

import (
	"math/big"
)

// ConcFactIter calculates factorial of the given
// number by splitting the task between workers
// for doing the work concurrently. At multi-core
// machines workers are making calculations in
// parallel.
//
// Utilizes slow backend for multiplication: mulRangeIter.
func ConcFactIter(n, workerNum uint64) *big.Int {
	return concFact(n, workerNum, mulRangeIter)
}

// ConcFactFast calculates factorial of the given
// number by splitting the task between workers
// for doing the work concurrently. At multi-core
// machines workers are making calculations in
// parallel.
//
// Utilizes fast backend for multiplication: mulRangeFast.
func ConcFactFast(n, workerNum uint64) *big.Int {
	return concFact(n, workerNum, mulRangeFast)
}

func concFact(n uint64, workerNum uint64, fn func(a, b uint64) *big.Int) *big.Int {
	if n < 3 {
		retval, _ := miniFact(n)
		return retval
	}

	// n > 2

	// No need for concurrency in such case
	if n/workerNum < 2 {
		return fn(1, n)
	}

	ranges := numToRanges(n, workerNum) // split work
	ch := make(chan *big.Int)
	done := make(chan bool)

	for _, rng := range ranges {
		a, b := rng[0], rng[1]
		go func() {
			ch <- fn(a, b)
			done <- true
		}()
	}

	go func() {
		for i := 0; i < len(ranges); i++ {
			<-done
		}
		close(ch)
	}()

	res := big.NewInt(1)
	for num := range ch {
		res.Mul(res, num)
	}
	return res
}
