package fastfact

import (
	"errors"
	"math/big"
)

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
