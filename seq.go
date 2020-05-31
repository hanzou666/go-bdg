package bdg

type DnaSeq struct {
	Seq []byte
}

type Coordinate struct {
	Start int64
	End   int64
}

func NewDnaSeqFromStr(dnaStr string) *DnaSeq {
	return &DnaSeq{Seq: []byte(dnaStr)}
}

func (d DnaSeq) MarshalText() ([]byte, error) {
	return d.Seq, nil
}

func (d *DnaSeq) Len() int64 {
	return int64(len(d.Seq))
}
