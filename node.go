package bdg

type Node struct {
	Id  int64   `json:"node_id"`
	Seq *DnaSeq `json:"sequence"`
	Len int64
}

func NewNode(nodeId int64, seq *DnaSeq) *Node {
	return &Node{
		Id:  nodeId,
		Seq: seq,
		Len: seq.Len(),
	}
}

func NewNodeFromString(nodeId int64, seq string) *Node {
	dnaSeq := NewDnaSeqFromStr(seq)
	return NewNode(nodeId, dnaSeq)
}

func (n *Node) fixNodeId(baseId int64) *Node {
	return NewNode(baseId+n.Id, n.Seq)
}

func getMaxIdNode(nodes []*Node) (maxNode *Node) {
	maxNode = nodes[0]
	for _, node := range nodes {
		if node.Id > maxNode.Id {
			maxNode = node
		}
	}
	return
}

func makeNodeIndex(nodes []*Node) (id2Node map[int64]*Node) {
	id2Node = map[int64]*Node{}
	for _, node := range nodes {
		id2Node[node.Id] = node
	}
	return
}
