package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/blin/cracking-coding-interview/pkg/chapter-1/isuniq"
)

func main() {
	isuniq.GenerateGraphviz = true

	isuniq.IsUniqueNoDataStructures("asdfghjkl;")

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

	for i, g := range isuniq.IsUniqueNoDataStructuresGraphs {
		fn := path.Join(dotDir, fmt.Sprintf("%03d.dot", i))
		err := ioutil.WriteFile(fn, []byte(g.String()), 0644)
		if err != nil {
			fmt.Printf("failed to write out a dot file: %v\n", err)
			os.Exit(2)
		}
	}
}
