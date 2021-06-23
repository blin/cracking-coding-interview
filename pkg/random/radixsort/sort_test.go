package radixsort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	xs := []int{170, 45, 75, 90, 2, 802, 2, 66}
	want := []int{2, 2, 45, 66, 75, 90, 170, 803}
	Sort(xs)
	if !reflect.DeepEqual(xs, want) {
		t.Errorf("got=%v, want=%v", xs, want)
	}
}
