package graph

type Graph struct {
	Nodes map[string]*Node
}

type Node struct {
	Addr  PropertyAddr
	Edges map[*Node]EdgeType

	// Mark this node as ignored when validating the graph.
	Ignored bool

	graph *Graph
}

type EdgeType int

const (
	ConflictEdge EdgeType = iota
	RequireEdge
)

func (g *Graph) AddNode(addr string) {
	node := Node{
		Addr:  MustParseAddr(addr),
		graph: g,
	}
	g.Nodes[node.Addr.String()] = &node
}

func (n *Node) AddEdges(addr PropertyAddr, t EdgeType) {
	g := n.graph
	for addrStr, dstNode := range g.Nodes {
		if dstNode == n {
			continue
		}
		nodeAddr := MustParseAddr(addrStr)
		if nodeAddr.Belongs(addr) {
			n.Edges[dstNode] = t
			dstNode.Edges[n] = t
		}
	}
}

type Solution struct {
	IgnoredAddrs map[string]bool
}

func (g *Graph) ValidSolution(solutions []Solution) *Solution {
	for _, solution := range solutions {
		stop := false
		func() {
			defer func() {
				// Reset nodes
				for _, node := range g.Nodes {
					node.Ignored = false
				}
			}()

			for iaddr := range solution.IgnoredAddrs {
				iaddr := MustParseAddr(iaddr)
				for addr, node := range g.Nodes {
					addr := MustParseAddr(addr)
					if addr.Belongs(iaddr) {
						node.Ignored = true
					}
				}
			}

			stop = g.Valid()
		}()
		if stop {
			return &solution
		}
	}
	// This should never happen
	return nil
}

func (g *Graph) Valid() bool {
	for _, node := range g.Nodes {
		if node.Ignored {
			continue
		}
		for onode, et := range node.Edges {
			switch et {
			case RequireEdge:
				if onode.Ignored {
					return false
				}
			case ConflictEdge:
				if !onode.Ignored {
					return false
				}
			}
		}
	}
	return true
}
