package bdg

type Graph struct {
	Nodes     []*Node `json:"node,omitempty"`
	Edges     []*Edge `json:"edge,omitempty"`
	Paths     []*Path `json:"path,omitempty"`
	nodeIndex map[int64]*Node
	pathIndex map[string]*Path
}

func NewGraph(nodes []*Node, edges []*Edge, paths []*Path) *Graph {
	return &Graph{
		Nodes:     nodes,
		Edges:     edges,
		Paths:     paths,
		nodeIndex: makeNodeIndex(nodes),
		pathIndex: makePathIndex(paths),
	}
}

func (g *Graph) fixNodeId(baseId int64) *Graph {
	var (
		newNodes []*Node
		newEdges []*Edge
		newPaths []*Path
	)
	for _, node := range g.Nodes {
		newNodes = append(newNodes, node.fixNodeId(baseId))
	}
	for _, edge := range g.Edges {
		newEdges = append(newEdges, edge.fixNodeId(baseId))
	}
	for _, path := range g.Paths {
		newPaths = append(newPaths, path.fixNodeId(baseId))
	}

	return NewGraph(newNodes, newEdges, newPaths)
}

func Combine(graphs []*Graph) *Graph {
	var newGraph Graph
	nextId := int64(0)
	for _, graph := range graphs {
		fixedGraph := graph.fixNodeId(nextId)
		newGraph.Nodes = append(newGraph.Nodes, fixedGraph.Nodes...)
		newGraph.Edges = append(newGraph.Edges, fixedGraph.Edges...)
		newGraph.Paths = append(newGraph.Paths, fixedGraph.Paths...)
		nextId = getMaxIdNode(newGraph.Nodes).Id
	}
	newGraph.nodeIndex = makeNodeIndex(newGraph.Nodes)
	newGraph.pathIndex = makePathIndex(newGraph.Paths)
	return &newGraph
}

func (g *Graph) GetPath(pathName string) (path *Path, ok bool) {
	path, ok = g.pathIndex[pathName]
	return
}
