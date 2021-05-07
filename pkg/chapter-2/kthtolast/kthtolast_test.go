package kthtolast

import (
	"testing"

	"github.com/blin/cracking-coding-interview/pkg/listutils"
)

func TestDups(t *testing.T) {
	cases := []struct {
		name      string
		l         []int
		k         int
		want      int
		wantError bool
	}{
		{
			name:      "get kth to last for empty list",
			l:         []int{},
			k:         0,
			wantError: true,
		},
		{
			name:      "get last element for single element list",
			l:         []int{0},
			k:         0,
			want:      0,
			wantError: false,
		},
		{
			name:      "get kth last element past list end",
			l:         []int{0},
			k:         1,
			wantError: true,
		},
		{
			name:      "get last element",
			l:         []int{6, 5, 4, 3, 2, 1, 0},
			k:         0,
			want:      0,
			wantError: false,
		},
		{
			name:      "get some kth to last",
			l:         []int{6, 5, 4, 3, 2, 1, 0},
			k:         3,
			want:      3,
			wantError: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			l := listutils.IntSliceToList(tc.l)
			got, err := KthToLast(l, tc.k)
			if (err != nil) != tc.wantError {
				t.Fatalf("got err=%v, wantError=%t", err, tc.wantError)
			}
			if err != nil {
				return
			}
			gotInt := int(got.(listutils.ComparableInt))
			if gotInt != tc.want {
				t.Errorf("got=%+v, want=%+v", gotInt, tc.want)
			}
		})
	}

}
