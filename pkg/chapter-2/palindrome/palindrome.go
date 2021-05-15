package palindrome

import (
	"container/list"

	"github.com/blin/cracking-coding-interview/pkg/listutils"
)

func IsPalindrome(l *list.List) bool {
	head := l.Front()
	tail := l.Back()
	for i := 0; i < l.Len()/2; i++ {
		headValue := head.Value.(listutils.HashEq)
		tailValue := tail.Value.(listutils.HashEq)
		if !headValue.Equal(tailValue) {
			return false
		}
		head = head.Next()
		tail = tail.Prev()
	}
	return true
}
