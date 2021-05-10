package partition

import (
	"testing"

	"github.com/blin/cracking-coding-interview/pkg/listutils"
)

func TestPartition(t *testing.T) {
	cases := []struct {
		name                     string
		l, wantBefore, wantAfter []int
	}{
		{
			name:       "list with one dup",
			l:          []int{0, 1, 2, 3, 4, 5},
			wantBefore: []int{0, 1, 2, 3},
			wantAfter:  []int{4, 5},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			l := listutils.IntSliceToList(tc.l)
			before, after := Partition(l, listutils.HashEqOrderedInt(4))
			if !listutils.ListEqual(before, listutils.IntSliceToList(tc.wantBefore)) {
				t.Errorf("got=%+v, want=%+v", listutils.ListToIntSlice(before), tc.wantBefore)
			}
			if !listutils.ListEqual(after, listutils.IntSliceToList(tc.wantAfter)) {
				t.Errorf("got=%+v, want=%+v", listutils.ListToIntSlice(after), tc.wantAfter)
			}
		})
	}
}
