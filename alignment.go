package bdg

type Alignment struct {
	Name     string      `json:"name"`
	Path     *Path       `json:"path"`
	RefPos   []*Position `json:"refpos,omitempty"`
	Sequence *DnaSeq     `json:"sequence,omitempty"`
}
