package dups

import (
	"container/list"

	"github.com/blin/cracking-coding-interview/pkg/listutils"
)

// Duplicates are taken to mean "list elements with the same value"
func RemoveDuplicates(l *list.List) {
	// TODO: collisions
	values := map[uint64][]listutils.HashEq{}
	e := l.Front()
	for {
		if e == nil {
			break
		}

		v, ok := e.Value.(listutils.HashEq)
		if !ok {
			panic("non comparable value in list")
		}

		foundCollision := false
		valuesWithSameHash := values[v.Hash()]
		for _, valueWithSameHash := range valuesWithSameHash {
			if v.Equal(valueWithSameHash) {
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
