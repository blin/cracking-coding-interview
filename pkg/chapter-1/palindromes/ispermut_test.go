package palindromes

import (
	"reflect"
	"testing"
)

func TestIsPalindromePermutation(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{
			s:    "ab ab",
			want: true,
		},
		{
			s:    "abc ab",
			want: true,
		},
		{
			s:    "abc abd",
			want: false,
		},
	}
	for _, tc := range cases {
		got := IsPalindromePermutation(tc.s)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf(`IsPalindromePermutation(%v)==%v, expected %v`, tc.s, got, tc.want)
		}
	}
}

func TestIsPalindromePermutationBitVec(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{
			s:    "ab ab",
			want: true,
		},
		{
			s:    "abc ab",
			want: true,
		},
		{
			s:    "abc abd",
			want: false,
		},
	}
	for _, tc := range cases {
		got := IsPalindromePermutationBitVec(tc.s)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf(`IsPalindromePermutationBitVec(%v)==%v, expected %v`, tc.s, got, tc.want)
		}
	}
}
