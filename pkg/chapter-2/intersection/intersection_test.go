package intersection

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	cases := []struct {
		name              string
		l1, l2            []int
		intersectionValue int
	}{
		{
			name:              "simple intersection",
			l1:                []int{10, 11, 12},
			l2:                []int{20, 21, 11},
			intersectionValue: 11,
		},
		{
			name:              "no intersection",
			l1:                []int{10, 11, 12},
			l2:                []int{20, 21, 11},
			intersectionValue: 999,
		},
	}

	for _, tc := range cases {
		// TODO: test that l1 and l2 look as expected
		t.Run(tc.name, func(t *testing.T) {
			head1, head2, intersectionElement := GenerateIntersectingLists(tc.l1, tc.l2, tc.intersectionValue)

			gotIntersectionElement := FindIntersection(head1, head2)

			if gotIntersectionElement != intersectionElement {
				t.Errorf("got=%+v, want=%+v", gotIntersectionElement, intersectionElement)
			}

		})
	}

}
