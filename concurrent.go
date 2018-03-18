package fastfact

import (
	"math/big"
)

// ConcFact calculates factorial of the given
// number by splitting the task between workers
// for doing the work concurrently. At multi-core
// machines workers are making calculations in
// parallel.
func ConcFact(n, workerNum uint64) *big.Int {
	if n < 3 {
		retval, _ := miniFact(n)
		return retval
	}

	// n > 2

	// No need for concurrency in such case
	if n/workerNum < 2 {
		return mulRange(1, n)
	}

	ranges := numToRanges(n, workerNum) // split work
	ch := make(chan *big.Int)
	done := make(chan bool)

	for _, rng := range ranges {
		a, b := rng[0], rng[1]
		go func() {
			ch <- mulRange(a, b)
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
