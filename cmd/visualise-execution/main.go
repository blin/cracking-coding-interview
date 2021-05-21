package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/blin/cracking-coding-interview/pkg/chapter-2/intersection"
)

func main() {
	intersection.GenerateGraphviz = true

	head1, head2, _ := intersection.GenerateIntersectingLists([]int{10, 11, 12}, []int{20, 21, 11}, 11)
	intersection.FindIntersection(head1, head2)

	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("failed to get working directory: %v\n", err)
		os.Exit(2)
	}
	dotDir := path.Join(wd, "dot")
	if _, err := os.Stat(dotDir); os.IsNotExist(err) {
		os.Mkdir(dotDir, 0755)
	} else if err != nil {
		fmt.Printf("failed to stat %s : %v\n", dotDir, err)
		os.Exit(2)
	}

	for i, g := range intersection.FindIntersectionGraphs {
		fn := path.Join(dotDir, fmt.Sprintf("%03d.dot", i))
		err := ioutil.WriteFile(fn, []byte(g.String()), 0644)
		if err != nil {
			fmt.Printf("failed to write out a dot file: %v\n", err)
			os.Exit(2)
		}
	}
}
