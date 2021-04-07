package matrices

import (
	"reflect"
	"testing"
)

func TestComputeSquarePosition(t *testing.T) {
	cases := []struct {
		sideSize int
		depth    int
		idx      int
		wantY    int
		wantX    int
	}{
		{
			sideSize: 6,
			depth:    0,
			idx:      0,
			wantY:    0,
			wantX:    0,
		},
		{
			sideSize: 6,
			depth:    0,
			idx:      5,
			wantY:    0,
			wantX:    5,
		},
		{
			sideSize: 6,
			depth:    0,
			idx:      6,
			wantY:    1,
			wantX:    5,
		},
		{
			sideSize: 6,
			depth:    0,
			idx:      10,
			wantY:    5,
			wantX:    5,
		},
		{
			sideSize: 6,
			depth:    0,
			idx:      11,
			wantY:    5,
			wantX:    4,
		},
		{
			sideSize: 6,
			depth:    0,
			idx:      15,
			wantY:    5,
			wantX:    0,
		},
		{
			sideSize: 6,
			depth:    0,
			idx:      16,
			wantY:    4,
			wantX:    0,
		},
		{
			sideSize: 6,
			depth:    0,
			idx:      19,
			wantY:    1,
			wantX:    0,
		},
		{
			sideSize: 6,
			depth:    1,
			idx:      0,
			wantY:    1,
			wantX:    1,
		},
		{
			sideSize: 6,
			depth:    1,
			idx:      4,
			wantY:    2,
			wantX:    4,
		},
		{
			sideSize: 6,
			depth:    1,
			idx:      7,
			wantY:    4,
			wantX:    3,
		},
		{
			sideSize: 6,
			depth:    1,
			idx:      10,
			wantY:    3,
			wantX:    1,
		},
	}

	for _, tc := range cases {
		gotY, gotX := computeSquarePosition(tc.sideSize, tc.depth, tc.idx)
		if gotY != tc.wantY || gotX != tc.wantX {
			t.Errorf("computeSquarePosition(%d, %d, %d)==(%d, %d), expected (%d, %d)", tc.sideSize, tc.depth, tc.idx, gotY, gotX, tc.wantY, tc.wantX)
		}
	}
}

func TestRotate(t *testing.T) {
	cases := []struct {
		yxM  [][]uint32
		want [][]uint32
	}{
		{
			yxM: [][]uint32{
				{1, 2, 3, 4, 5, 6},
				{20, 21, 22, 23, 24, 7},
				{19, 32, 33, 34, 25, 8},
				{18, 31, 36, 35, 26, 9},
				{17, 30, 29, 28, 27, 10},
				{16, 15, 14, 13, 12, 11},
			},
			want: [][]uint32{
				{16, 17, 18, 19, 20, 1},
				{15, 30, 31, 32, 21, 2},
				{14, 29, 36, 33, 22, 3},
				{13, 28, 35, 34, 23, 4},
				{12, 27, 26, 25, 24, 5},
				{11, 10, 9, 8, 7, 6},
			},
		},
		{
			yxM: [][]uint32{
				{1, 2, 3, 4, 5},
				{16, 17, 18, 19, 6},
				{15, 24, 25, 20, 7},
				{14, 23, 22, 21, 8},
				{13, 12, 11, 10, 9},
			},
			want: [][]uint32{
				{13, 14, 15, 16, 1},
				{12, 23, 24, 17, 2},
				{11, 22, 25, 18, 3},
				{10, 21, 20, 19, 4},
				{9, 8, 7, 6, 5},
			},
		},
	}

	for _, tc := range cases {
		Rotate(tc.yxM, nil)
		if !reflect.DeepEqual(tc.yxM, tc.want) {
			t.Errorf(`Rotate()==%v, expected %v`, tc.yxM, tc.want)

		}

	}
}
