package isuniq

import (
	"math/rand"
	"testing"
	"unicode"
)

func TestIsUniq(t *testing.T) {
	if !IsUnique("abc") {
		t.Errorf(`got IsUnique("abc")==false, expected true`)
	}
	if IsUnique("aaa") {
		t.Errorf(`got IsUnique("aaa")==true, expected false`)
	}
}

func TestIsUniqueNoDataStructures(t *testing.T) {
	if !IsUniqueNoDataStructures("abc") {
		t.Errorf(`got IsUniqueNoDataStructures("abc")==false, expected true`)
	}
	if IsUniqueNoDataStructures("aaa") {
		t.Errorf(`got IsUniqueNoDataStructures("aaa")==true, expected false`)
	}
}

const size int = 1000

var list = make([]rune, size)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringWithRepetition(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandStringWithoutRepetition(n int) string {
	r := rune(0)
	b := []rune{}
	for {
		r++
		if !unicode.IsLetter(r) {
			continue
		}
		b = append(b, r)

		if len(b) > n {
			break
		}
	}
	rand.Shuffle(len(b), func(i, j int) { b[i], b[j] = b[j], b[i] })
	return string(b)
}

func BenchmarkIsUnique(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := RandStringWithRepetition(size)
		b.StartTimer()
		IsUnique(s)
	}
}

func BenchmarkIsUniqueWorstCase(b *testing.B) {
	s := RandStringWithoutRepetition(size)
	for n := 0; n < b.N; n++ {
		IsUnique(s)
	}
}

func BenchmarkIsUniqueNoDataStructures(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := RandStringWithRepetition(size)
		b.StartTimer()
		IsUniqueNoDataStructures(s)
	}
}

func BenchmarkIsUniqueNoDataStructuresWorstCase(b *testing.B) {
	s := RandStringWithoutRepetition(size)
	for n := 0; n < b.N; n++ {
		IsUniqueNoDataStructures(s)
	}
}
