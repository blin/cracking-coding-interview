package urlify

import (
	"reflect"
	"testing"
)

func TestURLify(t *testing.T) {
	cases := []struct {
		s    []rune
		want []rune
	}{
		{
			s:    []rune("a a"),
			want: []rune("a%20a"),
		},
		{
			s:    []rune("    "),
			want: []rune("%20%20%20%20"),
		},
	}
	for _, tc := range cases {
		got := URLify(tc.s)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf(`URLify(%v)==%v, expected %v`, tc.s, got, tc.want)
		}
	}
}
