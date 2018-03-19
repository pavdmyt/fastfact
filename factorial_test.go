package fastfact

import (
	"math/big"
	"testing"
)

var testCases = []struct {
	in  uint64
	out *big.Int
}{
	{0, big.NewInt(1)},
	{1, big.NewInt(1)},
	{2, big.NewInt(2)},
	{3, big.NewInt(6)},
	{5, big.NewInt(120)},
	{8, big.NewInt(40320)},
	{10, big.NewInt(3628800)},
	{12, big.NewInt(479001600)},
	{25, MulRangeFact(25)},
}

func TestMulRangeFact(t *testing.T) {
	for _, test := range testCases {
		res := MulRangeFact(test.in)
		if res.Cmp(test.out) != 0 {
			t.Fatalf("found %v, want %v", res, test.out)
		}
	}
}

func TestIterFact(t *testing.T) {
	for _, test := range testCases {
		res := IterFact(test.in)
		if res.Cmp(test.out) != 0 {
			t.Fatalf("found %v, want %v", res, test.out)
		}
	}
}

func TestHalfIterFact(t *testing.T) {
	for _, test := range testCases {
		res := HalfIterFact(test.in)
		if res.Cmp(test.out) != 0 {
			t.Fatalf("found %v, want %v", res, test.out)
		}
	}
}

func TestConcFactIter(t *testing.T) {
	for numW := 1; numW <= 20; numW++ { // number of workers
		for _, test := range testCases {
			res := ConcFactIter(test.in, uint64(numW))
			if res.Cmp(test.out) != 0 {
				t.Fatalf("found %v, want %v; numW: %v", res, test.out, numW)
			}
		}
	}
}

func TestConcFactFast(t *testing.T) {
	for numW := 1; numW <= 20; numW++ { // number of workers
		for _, test := range testCases {
			res := ConcFactFast(test.in, uint64(numW))
			if res.Cmp(test.out) != 0 {
				t.Fatalf("found %v, want %v; numW: %v", res, test.out, numW)
			}
		}
	}
}
