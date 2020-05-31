package bdg

type Edge struct {
	FromId    int64 `json:"from"`
	ToId      int64 `json:"to"`
	FromStart bool  `json:"from_start,omitempty"`
	ToEnd     bool  `json:"to_end,omitempty"`
}

func (e *Edge) fixNodeId(baseId int64) *Edge {
	return &Edge{
		FromId:    baseId + e.FromId,
		ToId:      baseId + e.ToId,
		FromStart: e.FromStart,
		ToEnd:     e.ToEnd,
	}
}
