package palindrome

import (
	"testing"

	"github.com/blin/cracking-coding-interview/pkg/listutils"
)

func TestDups(t *testing.T) {
	cases := []struct {
		name string
		l    []int
		want bool
	}{
		{
			name: "even not a palindrome",
			l:    []int{1, 2, 3, 1},
			want: false,
		},
		{
			name: "odd not a palindrome",
			l:    []int{1, 2, 3},
			want: false,
		},
		{
			name: "even palindrome",
			l:    []int{1, 2, 2, 1},
			want: true,
		},
		{
			name: "odd palindrome",
			l:    []int{1, 2, 3, 2, 1},
			want: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			l := listutils.IntSliceToList(tc.l)
			got := IsPalindrome(l)
			if got != tc.want {
				t.Errorf("got=%t, want=%t", got, tc.want)
			}
		})
	}

}
