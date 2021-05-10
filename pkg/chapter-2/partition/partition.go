package partition

import (
	"container/list"

	"github.com/blin/cracking-coding-interview/pkg/listutils"
)

func Partition(l *list.List, pivot listutils.Ordered) (*list.List, *list.List) {
	if l == nil {
		panic("list to partition must not be nil")
	}
	if pivot == nil {
		panic("pivot must not be nil")
	}

	before := list.New()
	after := list.New()

	for {
		current := l.Front()
		if current == nil {
			break
		}
		currentValue := l.Remove(current)

		valueCompared := pivot.Compare(currentValue) * -1

		switch {
		case valueCompared < 0:
			before.PushBack(currentValue)
		case valueCompared >= 0:
			after.PushBack(currentValue)
		}
	}

	return before, after
}
