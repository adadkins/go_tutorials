package breadth_first_search_test

import (
	"go_tutorials/breadth_first_search"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFS(t *testing.T) {
	t.Run("BFS finds value", func(t *testing.T) {
		node1 := &breadth_first_search.Node{}
		node1.SetValue(1)
		node2 := &breadth_first_search.Node{}
		node2.SetValue(2)

		node3 := &breadth_first_search.Node{}
		node3.SetValue(3)

		node4 := &breadth_first_search.Node{}
		node4.SetValue(4)

		node1.AddNeighbor(node2)
		node1.AddNeighbor(node3)
		node2.AddNeighbor(node4)

		graph := &breadth_first_search.Graph{}
		graph.AddNode(node1)
		graph.AddNode(node2)
		graph.AddNode(node3)
		graph.AddNode(node4)

		result := graph.Breadth_first_search(node1, 4)

		assert.Equal(t, true, result)
	})

	t.Run("BFS finds value on larger graph", func(t *testing.T) {
		nodes := make([]*breadth_first_search.Node, 10)
		for i := range nodes {
			nodes[i] = &breadth_first_search.Node{}
			nodes[i].SetValue(i)
		}
		// create some neighbors
		for i := range nodes {
			nodes[i].AddNeighbor(nodes[(i+1)%10])
			nodes[i].AddNeighbor(nodes[(i+5)%10])
		}

		// add nodes to graph
		graph := &breadth_first_search.Graph{}
		for _, node := range nodes {
			graph.AddNode(node)
		}

		result := graph.Breadth_first_search(nodes[0], 5)
		assert.Equal(t, true, result)

	})

	t.Run("BFS doesnt find value", func(t *testing.T) {
		nodes := make([]*breadth_first_search.Node, 10)
		for i := range nodes {
			nodes[i] = &breadth_first_search.Node{}
			nodes[i].SetValue(i)
		}

		// create some neighbors
		for i := range nodes {
			nodes[i].AddNeighbor(nodes[(i+1)%10])
			nodes[i].AddNeighbor(nodes[(i+5)%10])
		}

		// add nodes to graph
		graph := &breadth_first_search.Graph{}
		for _, node := range nodes {
			graph.AddNode(node)
		}

		result := graph.Breadth_first_search(nodes[0], 500)
		assert.Equal(t, false, result)
	})
}
