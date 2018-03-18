package fastfact

import (
	"math/big"
	"reflect"
	"testing"
)

var numToRangesCases = []struct {
	num     uint64
	workers uint64
	out     [][]uint64
}{
	// num % workers == 0
	{
		10, 1, [][]uint64{
			{1, 10},
		},
	},
	{
		10, 2, [][]uint64{
			{1, 5}, {6, 10},
		},
	},
	{
		8, 4, [][]uint64{
			{1, 2}, {3, 4}, {5, 6}, {7, 8},
		},
	},
	{
		9, 3, [][]uint64{
			{1, 3}, {4, 6}, {7, 9},
		},
	},
	// num % workers == 1
	{
		10, 3, [][]uint64{
			{1, 3}, {4, 6}, {7, 10},
		},
	},
	// num % workers == 2
	{
		11, 3, [][]uint64{
			{1, 4}, {5, 8}, {9, 11},
		},
	},
	{
		17, 3, [][]uint64{
			{1, 6}, {7, 12}, {13, 17},
		},
	},
	// num % workers == 3
	{
		11, 4, [][]uint64{
			{1, 3}, {4, 6}, {7, 9}, {10, 11},
		},
	},
}

var mulRangeCases = []struct {
	a   uint64
	b   uint64
	out *big.Int
}{
	{2, 1, big.NewInt(1)}, // a > b
	{3, 4, big.NewInt(3 * 4)},
	{4, 6, big.NewInt(4 * 5 * 6)},
	{5, 8, big.NewInt(5 * 6 * 7 * 8)},
	{2, 9, big.NewInt(2 * 3 * 4 * 5 * 6 * 7 * 8 * 9)},
}

func TestNumToRanges(t *testing.T) {
	for _, test := range numToRangesCases {
		ranges := numToRanges(test.num, test.workers)
		if !reflect.DeepEqual(ranges, test.out) {
			t.Fatalf("%v is not equal to %v", ranges, test.out)
		}
	}
}

func TestMulRange(t *testing.T) {
	for _, test := range mulRangeCases {
		res := mulRange(test.a, test.b)
		if res.Cmp(test.out) != 0 {
			t.Fatalf("found %v, want %v", res, test.out)
		}
	}
}
