package matrices

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	yxM := [][]uint32{
		{1, 2, 3, 4, 5, 6},
		{20, 21, 22, 23, 24, 7},
		{19, 32, 33, 34, 25, 8},
		{18, 31, 36, 35, 26, 9},
		{17, 30, 29, 28, 27, 10},
		{16, 15, 14, 13, 12, 11},
	}

	want := [][]uint32{
		{16, 17, 18, 19, 20, 1},
		{15, 30, 31, 32, 21, 2},
		{14, 29, 36, 33, 22, 3},
		{13, 28, 35, 34, 23, 4},
		{12, 27, 26, 25, 24, 5},
		{11, 10, 9, 8, 7, 6},
	}
	// TODO: test odd size case
	Rotate(yxM, nil)
	if !reflect.DeepEqual(yxM, want) {
		t.Errorf(`Rotate()==%v, expected %v`, yxM, want)

	}
}
