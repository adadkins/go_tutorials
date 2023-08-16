package breadth_first_search

type Graph struct {
	nodes []*Node
}

type Node struct {
	neighbors []*Node
	value     int
	visited   bool
}

func (n *Node) SetValue(input int) {
	n.value = input
}

func (n *Node) AddNeighbor(neighbor *Node) {
	n.neighbors = append(n.neighbors, neighbor)
}

func (g *Graph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) Breadth_first_search(start *Node, target int) bool {
	queue := make([]*Node, 0)
	start.visited = true
	queue = append(queue, start)

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		if currentNode.value == target {
			return true
		}
		for _, neighbor := range currentNode.neighbors {
			if !neighbor.visited {
				neighbor.visited = true
				queue = append(queue, neighbor)
			}
		}
	}

	return false
}
