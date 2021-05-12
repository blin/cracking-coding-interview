package sum

import (
	"testing"

	"github.com/blin/cracking-coding-interview/pkg/listutils"
)

func TestLittleEndianToInt(t *testing.T) {
	cases := []struct {
		name string
		l1   []int
		want int
	}{
		{
			name: "simplest case",
			l1:   []int{1, 2, 3},
			want: 321,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			l1 := listutils.IntSliceToList(tc.l1)
			got := BigEndianToInt(l1)
			if got != tc.want {
				t.Errorf("got=%+v, want=%+v", got, tc.want)
			}
		})
	}

}

func TestIntToLittleEndian(t *testing.T) {
	cases := []struct {
		name string
		x    int
		want []int
	}{
		{
			name: "simplest case",
			want: []int{1, 2, 3},
			x:    321,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			want := listutils.IntSliceToList(tc.want)
			got := IntToBigEndian(tc.x)
			if listutils.ListEqual(got, want) {
				t.Errorf("got=%+v, want=%+v", listutils.ListToIntSlice(got), listutils.ListToIntSlice(want))
			}
		})
	}

}
