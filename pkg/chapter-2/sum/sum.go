package sum

import (
	"container/list"
	"math"

	"github.com/blin/cracking-coding-interview/pkg/listutils"
)

func SumBigEndian(l1, l2 *list.List) *list.List {
	l1AsInt := BigEndianToInt(l1)
	l2AsInt := BigEndianToInt(l2)
	return IntToBigEndian(l1AsInt + l2AsInt)

}

func BigEndianToInt(l *list.List) int {
	var sum int

	i := 0
	for current := l.Front(); current != nil; current = current.Next() {
		elementValue := current.Value.(listutils.Int).Int()
		elementAsInt := elementValue * int(math.Pow(10, float64(i)))
		sum += elementAsInt
		i++
	}

	return sum
}

func IntToBigEndian(x int) *list.List {
	l := list.New()

	for current := x; current != 0; current = current / 10 {
		digit := current % 10
		l.PushFront(listutils.HashEqOrderedInt(digit))
	}

	return l
}
