package cycledetect

import "testing"

func TestDetectCycle(t *testing.T) {
	l := generateCyclicList([]int{0, 1, 2, 3, 4, 5})
	want := 0

	cycleElement, hasCycle := DetectCycle(l)
	if !hasCycle {
		t.Fatalf("didn't detect a cycle")
	}

	got := cycleElement.Value.(int)
	if got != want {
		t.Fatalf("got %d, want=%d", got, want)
	}
}
