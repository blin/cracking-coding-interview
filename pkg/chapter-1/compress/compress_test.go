package compress

import (
	"reflect"
	"testing"
)

func TestCompress(t *testing.T) {
	cases := []struct {
		s    string
		want string
	}{
		{
			s:    "ab",
			want: "ab",
		},
		{
			s:    "aaa",
			want: "a3",
		},
		{
			s:    "aaabbb",
			want: "a3b3",
		},
		{
			s:    "aaabbbc",
			want: "a3b3c1",
		},
	}
	for _, tc := range cases {
		got := Compress(tc.s)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf(`Compress(%v)==%v, expected %v`, tc.s, got, tc.want)
		}
	}
}
