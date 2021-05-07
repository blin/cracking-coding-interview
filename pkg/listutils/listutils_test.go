package listutils

import "testing"

func TestListEqual(t *testing.T) {
	cases := []struct {
		name   string
		l1, l2 []int
		want   bool
	}{
		{
			name: "simiple equal lists",
			l1:   []int{1, 2, 3},
			l2:   []int{1, 2, 3},
			want: true,
		},
		{
			name: "simiple unequal lists",
			l1:   []int{1, 2, 3},
			l2:   []int{1, 2},
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ListEqual(IntSliceToList(tc.l1), IntSliceToList(tc.l2))
			if got != tc.want {
				t.Errorf("got=%t, want=%t", got, tc.want)
			}
		})
	}
}
