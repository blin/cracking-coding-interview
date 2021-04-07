package isrotat

import "testing"

func TestIsRotation(t *testing.T) {
	cases := []struct {
		s1, s2 string
		want   bool
	}{
		{
			s1:   "aaa",
			s2:   "aaa",
			want: true,
		},
		{
			s1:   "aaa",
			s2:   "abc",
			want: false,
		},
		{
			s1:   "aaa",
			s2:   "aaaa",
			want: false,
		},
		{
			s1:   "abc",
			s2:   "bca",
			want: true,
		},
	}
	for _, tc := range cases {
		got := IsRotation(tc.s1, tc.s2)
		if got != tc.want {
			t.Errorf(`IsRotation("%s", "%s")==%t, expected %t`, tc.s1, tc.s2, got, tc.want)
		}
	}
}
