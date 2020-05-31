package bdg

type Position struct {
	NodeId     int64  `json:"node_id"`
	Offset     int64  `json:"offset"`
	IsReversed bool   `json:"is_reverse,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Edit struct {
	FromLength int32   `json:"from_length,omitempty"`
	ToLength   int32   `json:"to_length,omitempty"`
	Sequence   *DnaSeq `json:"sequence,omitempty"`
}

type Mapping struct {
	Position *Position `json:"position"`
	Edits    []*Edit   `json:"edit"`
	Rank     int64     `json:"rank"`
}

type Path struct {
	Name     string     `json:"name"`
	Mappings []*Mapping `json:"mapping"`
}

func MakeEditsFromNode(n *Node) []*Edit {
	return []*Edit{{
		FromLength: int32(n.Len),
		ToLength:   int32(n.Len),
		Sequence:   n.Seq,
	}}
}

func (p *Position) fixNodeId(baseId int64) *Position {
	return &Position{
		NodeId:     baseId + p.NodeId,
		Offset:     p.Offset,
		IsReversed: p.IsReversed,
		Name:       p.Name,
	}
}

func (m *Mapping) fixNodeId(baseId int64) *Mapping {
	return &Mapping{
		Position: m.Position.fixNodeId(baseId),
		Edits:    m.Edits,
		Rank:     m.Rank,
	}
}

func (p *Path) fixNodeId(baseId int64) *Path {
	var newMappings []*Mapping
	for _, mapping := range p.Mappings {
		newMappings = append(newMappings, mapping.fixNodeId(baseId))
	}
	return &Path{
		Name:     p.Name,
		Mappings: newMappings,
	}
}
