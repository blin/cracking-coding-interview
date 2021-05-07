package dups

import (
	"container/list"
)

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
