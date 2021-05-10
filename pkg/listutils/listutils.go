package listutils

import (
	"container/list"
	"reflect"
)

type HashEq interface {
	Hash() uint64
	Equal(interface{}) bool
}

type HashEqInt int

func (i HashEqInt) Hash() uint64 {
	return uint64(i)
}

func (i HashEqInt) Equal(o interface{}) bool {
	return reflect.DeepEqual(i, o)
}

func IntSliceToList(s []int) *list.List {
	l := list.New()

	for _, e := range s {
		l.PushBack(HashEqInt(e))
	}

	return l
}

func ListToIntSlice(l *list.List) []int {
	var s []int

	for e := l.Front(); e != nil; e = e.Next() {
		v, ok := e.Value.(HashEqInt)
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
