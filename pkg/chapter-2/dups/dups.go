package dups

import (
	"container/list"
	"reflect"
)

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

type Comparable interface {
	Hash() uint64
	Compare(interface{}) bool
}

// Duplicates are taken to mean "list elements with the same value"
func RemoveDuplicates(l *list.List) {
	// TODO: collisions
	values := map[uint64][]Comparable{}
	e := l.Front()
	for {
		if e == nil {
			break
		}

		v, ok := e.Value.(Comparable)
		if !ok {
			panic("non comparable value in list")
		}

		foundCollision := false
		valuesWithSameHash := values[v.Hash()]
		for _, valueWithSameHash := range valuesWithSameHash {
			if v.Compare(valueWithSameHash) {
				foundCollision = true
			}
		}
		if !foundCollision {
			values[v.Hash()] = append(values[v.Hash()], v)
			e = e.Next()
			continue
		}

		nextE := e.Next()
		l.Remove(e)
		e = nextE
	}
}
