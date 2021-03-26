package isuniq

import (
	"fmt"

	"github.com/awalterschulze/gographviz"
)

var GenerateGraphviz = false

var IsUniqueGraphs = []*gographviz.Graph{}

func charIdxToNodeName(charIdx int) string {
	return fmt.Sprintf("c%dk", charIdx)
}

func charToNodeName(c rune, charIndexes map[rune]int) string {
	return charIdxToNodeName(charIndexes[c])
}

func addCharSet(g *gographviz.Graph, charSet map[rune]bool, charIndexes map[rune]int) {
	g.AddSubGraph("G", "charSet", nil)
	for c := range charSet {
		g.AddNode("charSet", charToNodeName(c, charIndexes), map[string]string{
			"label": fmt.Sprintf(`"%c"`, c),
		})

		charIdx := charIndexes[c]
		if charIdx > 0 {
			srcNode := charIdxToNodeName(charIdx - 1)
			dstNode := charIdxToNodeName(charIdx)
			g.AddEdge(srcNode, dstNode, true, nil)
		}
	}
}

func IsUnique(s string) bool {
	if GenerateGraphviz {
		g := GenerateStringBaseGraph(s)

		IsUniqueGraphs = append(IsUniqueNoDataStructuresGraphs, g)
	}
	charSet := make(map[rune]bool, len(s))
	for i, r := range s {
		var g *gographviz.Graph
		charIndexes := map[rune]int{}
		if GenerateGraphviz {
			charIdx := 0
			for c := range charSet {
				charIndexes[c] = charIdx
				charIdx++
			}

			g = GenerateStringBaseGraph(s)

			n1 := StringIndexToNode(g, i)
			n1.Attrs[gographviz.Style] = "filled"
			n1.Attrs[gographviz.FillColor] = "blue"

			addCharSet(g, charSet, charIndexes)

			IsUniqueGraphs = append(IsUniqueGraphs, g)
		}
		if charSet[r] {
			if g != nil {
				kn := g.Nodes.Lookup[charToNodeName(r, charIndexes)]
				kn.Attrs[gographviz.Style] = "filled"
				kn.Attrs[gographviz.FillColor] = "red"
			}
			return false
		}
		charSet[r] = true
	}
	return true
}

var IsUniqueNoDataStructuresGraphs = []*gographviz.Graph{}

// TODO: make private, here and elsewhere
func StringIndexToNodeName(i int) string {
	return fmt.Sprintf("s%d", i)
}

func StringIndexToNode(g *gographviz.Graph, i int) *gographviz.Node {
	return g.Nodes.Lookup[StringIndexToNodeName(i)]
}

func GenerateStringBaseGraph(s string) *gographviz.Graph {
	g := gographviz.NewGraph()
	g.SetName("G")
	g.SetDir(true)
	g.AddAttr("G", "rankdir", "LR")
	for i, r := range s {
		g.AddNode("G", StringIndexToNodeName(i), map[string]string{
			"label": fmt.Sprintf(`"%c"`, r),
		})

		if i == 0 {
			continue
		}
		srcNode := StringIndexToNodeName(i - 1)
		dstNode := StringIndexToNodeName(i)
		g.AddEdge(srcNode, dstNode, true, nil)
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
