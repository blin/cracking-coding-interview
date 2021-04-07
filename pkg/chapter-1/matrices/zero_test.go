package matrices

import (
	"reflect"
	"testing"
)

func TestZeroCrossForZeroElement(t *testing.T) {
	cases := []struct {
		yxM  [][]uint32
		want [][]uint32
	}{
		{
			yxM: [][]uint32{
				{0, 2, 3, 4, 5, 6},
				{20, 21, 22, 23, 24, 7},
				{19, 32, 33, 34, 25, 8},
				{18, 31, 36, 35, 26, 9},
				{17, 30, 29, 28, 27, 10},
				{16, 15, 14, 13, 12, 11},
			},
			want: [][]uint32{
				{0, 0, 0, 0, 0, 0},
				{0, 21, 22, 23, 24, 7},
				{0, 32, 33, 34, 25, 8},
				{0, 31, 36, 35, 26, 9},
				{0, 30, 29, 28, 27, 10},
				{0, 15, 14, 13, 12, 11},
			},
		},
		{
			yxM: [][]uint32{
				{1, 2, 3, 4, 5, 6},
				{20, 21, 0, 23, 24, 7},
			},
			want: [][]uint32{
				{1, 2, 0, 4, 5, 6},
				{0, 0, 0, 0, 0, 0},
			},
		},
	}

	for _, tc := range cases {
		ZeroCrossForZeroElement(tc.yxM)
		if !reflect.DeepEqual(tc.yxM, tc.want) {
			t.Errorf(`Rotate()==%v, expected %v`, tc.yxM, tc.want)

		}

	}

}
