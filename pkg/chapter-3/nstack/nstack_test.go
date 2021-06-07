package nstack

import (
	"reflect"
	"testing"
)

func TestNStack(t *testing.T) {
	stacksData := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	want := [][]int{
		{3, 2, 1},
		{6, 5, 4},
		{9, 8, 7},
	}

	ns := New(len(stacksData))
	for i, stackData := range stacksData {
		for _, value := range stackData {
			ns.Push(i, value)
		}
	}

	var got [][]int
	for i := 0; i < len(stacksData); i++ {
		var s []int
		for {
			value, err := ns.Pop(i)
			if err == ErrIsEmpty {
				break
			}
			s = append(s, value.(int))
		}
		got = append(got, s)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got=%v, want=%v", got, want)
	}
}
