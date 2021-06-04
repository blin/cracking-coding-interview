package slist

type Element struct {
	Next *Element

	Value interface{}
}

func PushBack(head *Element, e *Element) {
	back := Back(head)
	back.Next = e
}

func Back(head *Element) *Element {
	current := head
	for current.Next != nil {
		current = current.Next
	}
	return current
}

