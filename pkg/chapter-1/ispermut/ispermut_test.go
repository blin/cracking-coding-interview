package ispermut

import "testing"

func TestIsPermutation(t *testing.T) {
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
			s2:   "asd",
			want: false,
		},
		{
			s1:   "aaa",
			s2:   "aaaa",
			want: false,
		},
	}
	for _, tc := range cases {
		got := IsPermutation(tc.s1, tc.s2)
		if got != tc.want {
			t.Errorf(`IsPermutation("%s", "%s")==%t, expected %t`, tc.s1, tc.s2, got, tc.want)
		}
	}
}

func TestIsPermutationNoAlloc(t *testing.T) {
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
			s2:   "asd",
			want: false,
		},
		{
			s1:   "aaa",
			s2:   "aaaa",
			want: false,
		},
	}
	for _, tc := range cases {
		got := IsPermutationNoAlloc(tc.s1, tc.s2)
		if got != tc.want {
			t.Errorf(`IsPermutationNoAlloc("%s", "%s")==%t, expected %t`, tc.s1, tc.s2, got, tc.want)
		}
	}
}
