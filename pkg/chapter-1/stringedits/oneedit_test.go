package stringedits

import "testing"

func TestIsOneEditAway(t *testing.T) {
	cases := []struct {
		s1, s2 string
		want   bool
	}{
		{
			s1:   "aaa",
			s2:   "aab",
			want: true,
		},
		{
			s1:   "aaa",
			s2:   "aaa",
			want: false,
		},
		{
			s1:   "aaa",
			s2:   "aaaa",
			want: true,
		},
		{
			s1:   "aaa",
			s2:   "abaa",
			want: true,
		},
		{
			s1:   "aaa",
			s2:   "baaa",
			want: true,
		},
		{
			s1:   "aaa",
			s2:   "bbaa",
			want: false,
		},
		{
			s1:   "aaaa",
			s2:   "aaa",
			want: true,
		},
		{
			s1:   "baaa",
			s2:   "aaa",
			want: true,
		},
	}
	for _, tc := range cases {
		got := IsOneEditAway(tc.s1, tc.s2)
		if got != tc.want {
			t.Errorf(`IsOneEditAway("%s", "%s")==%t, expected %t`, tc.s1, tc.s2, got, tc.want)
		}
	}
}
