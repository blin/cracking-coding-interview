package dups

import (
	"testing"

	"github.com/blin/cracking-coding-interview/pkg/listutils"
)

func TestDups(t *testing.T) {
	cases := []struct {
		name    string
		l, want []int
	}{
		{
			name: "list with one dup",
			l:    []int{1, 2, 1, 3},
			want: []int{1, 2, 3},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			l := listutils.IntSliceToList(tc.l)
			RemoveDuplicates(l)
			if !listutils.ListEqual(l, listutils.IntSliceToList(tc.want)) {
				t.Errorf("got=%+v, want=%+v", listutils.ListToIntSlice(l), tc.want)
			}
		})
	}

}
