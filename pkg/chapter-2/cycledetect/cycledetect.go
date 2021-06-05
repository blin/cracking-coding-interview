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

	// If the loop is of size n,
	// after i steps
	// p1 will be at position i%n
	// p2 will be at position d+(i*2))%n
	// where d is the distance between p1 and p2 when both enter the loop.
	//
	// Smallest i for which p1 == p2 is such that i == (n * (l1+l2)) - d
	// where l1 and l2 are number of loops completed by p1 and p2 respectively.
	//
	// Since p1 advances slower than p2, l1 = 0 and l2 = 1 is the smallest combination of loops
	// for which i == (n * (l1+l2)) - d holds true,
	// and so smallest i == n - d .
	//
	// Given the above, time to find the loop is bound by O(n).
	//
	// Thanks to Daria Zarayskaya
	// for helping me figure out the modular arithmetic for this algorithm.
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
