package listutils

import (
	"container/list"
	"reflect"
)

type Comparable interface {
	Hash() uint64
	Compare(interface{}) bool
}

type ComparableInt int

func (i ComparableInt) Hash() uint64 {
	return uint64(i)
}

func (i ComparableInt) Compare(o interface{}) bool {
	return reflect.DeepEqual(i, o)
}

func IntSliceToList(s []int) *list.List {
	l := list.New()

	for _, e := range s {
		l.PushBack(ComparableInt(e))
	}

	return l
}

func ListToIntSlice(l *list.List) []int {
	var s []int

	for e := l.Front(); e != nil; e = e.Next() {
		v, ok := e.Value.(ComparableInt)
		if !ok {
			panic("non int value in list")
		}
		s = append(s, int(v))
	}

	return s
}

func ListEqual(l, o *list.List) bool {
	if l.Len() != o.Len() {
		return false
	}
	lElem := l.Front()
	oElem := o.Front()
	for {
		if lElem == nil && oElem == nil {
			return true
		}
		if !reflect.DeepEqual(lElem.Value, oElem.Value) {
			return false
		}
		lElem = lElem.Next()
		oElem = oElem.Next()
	}
}
