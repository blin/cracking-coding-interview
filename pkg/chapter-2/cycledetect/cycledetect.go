package cycledetect

import (
	"github.com/blin/cracking-coding-interview/pkg/slist"
)

func hasDoubleNext(l *slist.Element) bool {
	return l.Next != nil && l.Next.Next != nil
}

func DetectCycle(l *slist.Element) (*slist.Element, bool) {
	p1 := l

	if !hasDoubleNext(l) {
		return nil, false
	}
	p2 := l.Next.Next

	// if the loop is of size n
	// p1 has a sequence [(i, i%n) for i in Z]
	// p2 has a sequence [(i, (k+(i*2))%n) for i in Z]
	// smallest i for which p1 == p2 seems to be n-k
	detectionIndex := 0
	for {
		if p1 == p2 {
			break
		}

		if !hasDoubleNext(p2) {
			return nil, false
		}

		p1 = p1.Next
		p2 = p2.Next.Next
		detectionIndex++
	}

	p1 = l
outer:
	for i := 0; i <= detectionIndex; i++ {
		p2 = p1.Next
		for j := 0; j <= (detectionIndex * 2); j++ {
			if p1 == p2 {
				break outer
			}
			p2 = p2.Next
		}
		p1 = p1.Next
	}
	return p1, true
}

func generateCyclicList(l []int) *slist.Element {
	var head *slist.Element
	for _, v := range l {
		e := &slist.Element{Value: v}
		if head == nil {
			head = e
			continue
		}
		slist.PushBack(head, e)
	}
	back := slist.Back(head)
	back.Next = head

	return head
}
