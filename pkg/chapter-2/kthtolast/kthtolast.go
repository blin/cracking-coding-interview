package kthtolast

import (
	"container/list"
	"fmt"
)

func KthToLast(l *list.List, k int) (interface{}, error) {
	if l.Len() == 0 {
		return nil, fmt.Errorf("there is no last element in list")
	}
	if l.Len() <= k {
		return nil, fmt.Errorf("list has len=%d, requested %d to last element", l.Len(), k)
	}

	p1 := l.Front()
	for i := 0; i < k; i++ {
		p1 = p1.Next()
	}

	p2 := l.Front()
	for {
		if p1.Next() == nil {
			return p2.Value, nil
		}
		p1 = p1.Next()
		p2 = p2.Next()
	}
}
