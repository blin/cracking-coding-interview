package intersection

import (
	"fmt"

	"github.com/awalterschulze/gographviz"
	"github.com/blin/cracking-coding-interview/pkg/slist"
)

var GenerateGraphviz = false

var FindIntersectionGraphs = []*gographviz.Graph{}

func generateBaseGraph(head1, head2 *slist.Element) *gographviz.Graph {
	g := gographviz.NewGraph()
	g.SetName("G")
	g.SetDir(true)
	g.AddAttr("G", "rankdir", "LR")

	addListToGraph(g, head1)
	addListToGraph(g, head2)

	return g
}

func recordBaseGraph(head1, head2 *slist.Element) {
	if !GenerateGraphviz {
		return
	}
	g := generateBaseGraph(head1, head2)

	FindIntersectionGraphs = append(FindIntersectionGraphs, g)
}

func recordElementProcessed(head1, head2, current1 *slist.Element, colour string) {
	if !GenerateGraphviz {
		return
	}
	g := generateBaseGraph(head1, head2)

	current1Node := elementToNode(g, current1)
	current1Node.Attrs[gographviz.Style] = "filled"
	current1Node.Attrs[gographviz.FillColor] = colour

	FindIntersectionGraphs = append(FindIntersectionGraphs, g)
}

func FindIntersection(head1, head2 *slist.Element) *slist.Element {
	recordBaseGraph(head1, head2)

	elementPointers := map[*slist.Element]bool{}
	current1 := head1
	for current1 != nil {
		recordElementProcessed(head1, head2, current1, "blue")

		elementPointers[current1] = true

		current1 = current1.Next
	}

	current2 := head2
	for current2 != nil {
		recordElementProcessed(head1, head2, current2, "blue")
		if _, present := elementPointers[current2]; present {
			recordElementProcessed(head1, head2, current2, "red")
			return current2
		}

		current2 = current2.Next
	}
	return nil
}

func elementToNode(g *gographviz.Graph, e *slist.Element) *gographviz.Node {
	eName := elementToNodeName(e)
	return g.Nodes.Lookup[eName]
}

func elementToNodeName(e *slist.Element) string {
	return fmt.Sprintf(`"%p"`, e)
}

func addListToGraph(g *gographviz.Graph, head *slist.Element) {
	current := head
	for current.Next != nil {
		currentName := elementToNodeName(current)
		nextName := elementToNodeName(current.Next)
		g.AddNode("G", currentName, map[string]string{
			"label": fmt.Sprintf(`"%v"`, current.Value),
		})
		g.AddNode("G", nextName, map[string]string{
			"label": fmt.Sprintf(`"%v"`, current.Next.Value),
		})
		g.AddEdge(currentName, nextName, true, nil)

		current = current.Next
	}
}

func GenerateIntersectingLists(l1, l2 []int, intersectionValue int) (*slist.Element, *slist.Element, *slist.Element) {
	var intersectionElement *slist.Element

	var head1, head2 *slist.Element

	for _, l1value := range l1 {
		e := &slist.Element{Value: l1value}
		if l1value == intersectionValue {
			intersectionElement = e
		}
		if head1 == nil {
			head1 = e
			continue
		}
		slist.PushBack(head1, e)
	}

	for _, l2value := range l2 {
		e := &slist.Element{Value: l2value}
		if l2value == intersectionValue {
			slist.Back(head2).Next = intersectionElement
			break
		}
		if head2 == nil {
			head2 = e
			continue
		}
		slist.PushBack(head2, e)
	}

	return head1, head2, intersectionElement
}
