package isuniq

import (
	"fmt"
	"sort"

	"github.com/awalterschulze/gographviz"
)

var GenerateGraphviz = false

func IsUnique(s string) bool {
	chars := make(map[rune]int, len(s))
	for _, r := range s {
		if chars[r] == 1 {
			return false
		}
		chars[r]++
	}
	return true
}

var IsUniqueNoDataStructuresGraphs = []*gographviz.Graph{}

func StringIndexToNodeName(i int) string {
	return fmt.Sprintf("s%d", i)
}

func StringIndexToNode(g *gographviz.Graph, i int) *gographviz.Node {
	return g.Nodes.Lookup[StringIndexToNodeName(i)]
}

func GenerateStringBaseGraph(s string) *gographviz.Graph {
	g := gographviz.NewGraph()
	g.SetName("G")
	for i, r := range s {
		g.AddNode("G", StringIndexToNodeName(i), map[string]string{
			"label": fmt.Sprintf(`"%c"`, r),
		})

		if i == 0 {
			continue
		}
		srcNode := StringIndexToNodeName(i - 1)
		dstNode := StringIndexToNodeName(i)
		g.AddEdge(srcNode, dstNode, false, nil)
	}
	return g
}

func IsUniqueNoDataStructures(s string) bool {
	if GenerateGraphviz {
		g := GenerateStringBaseGraph(s)

		IsUniqueNoDataStructuresGraphs = append(IsUniqueNoDataStructuresGraphs, g)
	}
	for i1, r1 := range s[:len(s)-1] {
		i2Offset := i1 + 1
		for i2, r2 := range s[i2Offset:] {
			if GenerateGraphviz {
				g := GenerateStringBaseGraph(s)

				n1 := StringIndexToNode(g, i1)
				n1.Attrs[gographviz.Style] = "filled"
				n1.Attrs[gographviz.FillColor] = "blue"

				n2 := StringIndexToNode(g, i2+i2Offset)
				n2.Attrs[gographviz.Style] = "filled"
				n2.Attrs[gographviz.FillColor] = "red"

				IsUniqueNoDataStructuresGraphs = append(IsUniqueNoDataStructuresGraphs, g)
			}
			if r1 == r2 {
				return false
			}
		}
	}
	return true
}

func IsUniqueWithSort(s string) bool {
	// []byte(string) -> runtime.stringtoslicebyte which allocates
	rs := []byte(s)
	// calling `func f(x interface{})` as `f([]byte)` results in runtime.convTslice call, which allocates
	// sort.Slice calls reflectlite.Swapper which allocates
	sort.Slice(rs, func(i, j int) bool { return rs[i] < rs[j] })

	// TODO test on empty string and one element string
	for i := 0; i < len(rs)-1; i++ {
		if rs[i] == rs[i+1] {
			return false
		}
	}

	return true
}
