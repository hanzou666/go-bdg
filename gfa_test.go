package bdg

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestEdge_linkLine(t *testing.T) {
	type fields struct {
		FromId    int64
		ToId      int64
		FromStart bool
		ToEnd     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "normal case",
			fields: fields{
				FromId:    1,
				ToId:      2,
				FromStart: false,
				ToEnd:     false,
			},
			want: "L\t1\t+\t2\t+\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Edge{
				FromId:    tt.fields.FromId,
				ToId:      tt.fields.ToId,
				FromStart: tt.fields.FromStart,
				ToEnd:     tt.fields.ToEnd,
			}
			if got := e.linkLine(); got != tt.want {
				t.Errorf("linkLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEdit_cigarStr(t *testing.T) {
	type fields struct {
		FromLength int32
		ToLength   int32
		Sequence   *DnaSeq
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "normal case: match",
			fields: fields{
				FromLength: 10,
				ToLength:   10,
				Sequence:   nil,
			},
			want: "10M",
		},
		{
			name: "normal case: substitution",
			fields: fields{
				FromLength: 10,
				ToLength:   10,
				Sequence:   NewDnaSeqFromStr("ATGCATGCAT"),
			},
			want: "10M",
		},
		{
			name: "normal case: deletion",
			fields: fields{
				FromLength: 1,
				ToLength:   0,
				Sequence:   nil,
			},
			want: "1D",
		},
		{
			name: "normal case: insertion",
			fields: fields{
				FromLength: 0,
				ToLength:   1,
				Sequence:   NewDnaSeqFromStr("A"),
			},
			want: "1I",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Edit{
				FromLength: tt.fields.FromLength,
				ToLength:   tt.fields.ToLength,
				Sequence:   tt.fields.Sequence,
			}
			if got := e.cigarStr(); got != tt.want {
				t.Errorf("cigarStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_WriteGFA(t *testing.T) {
	type fields struct {
		Nodes     []*Node
		Edges     []*Edge
		Paths     []*Path
		nodeIndex map[int64]*Node
		pathIndex map[string]*Path
	}
	type args struct {
		filename string
	}
	dir, err := ioutil.TempDir("./", "tmp_gfa")
	defer os.RemoveAll(dir)
	if err != nil {
		t.Error(err)
	}
	filename := filepath.Join(dir, "out.gfa")
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal case",
			fields: fields{
				Nodes: []*Node{
					NewNodeFromString(1, "ATCGA"),
					NewNodeFromString(2, "TGC"),
				},
				Edges: []*Edge{
					{
						FromId:    1,
						ToId:      2,
						FromStart: false,
						ToEnd:     false,
					},
				},
				Paths: []*Path{
					{
						Name: "chr1",
						Mappings: []*Mapping{
							{
								Position: &Position{
									NodeId:     1,
									Offset:     0,
									IsReversed: false,
									Name:       "",
								},
								Edits: []*Edit{
									{
										FromLength: 5,
										ToLength:   5,
										Sequence:   nil,
									},
								},
								Rank: 1,
							},
							{
								Position: &Position{
									NodeId:     2,
									Offset:     0,
									IsReversed: false,
									Name:       "",
								},
								Edits: []*Edit{
									{
										FromLength: 3,
										ToLength:   3,
										Sequence:   nil,
									},
								},
								Rank: 2,
							},
						},
					},
				},
				nodeIndex: nil,
				pathIndex: nil,
			},
			args:    args{filename},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Nodes:     tt.fields.Nodes,
				Edges:     tt.fields.Edges,
				Paths:     tt.fields.Paths,
				nodeIndex: tt.fields.nodeIndex,
				pathIndex: tt.fields.pathIndex,
			}
			if err := g.WriteGFA(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("WriteGFA() error = %v, wantErr %v", err, tt.wantErr)
			} else {

			}
		})
	}
}

func TestMapping_overlapStr(t *testing.T) {
	type fields struct {
		Position *Position
		Edits    []*Edit
		Rank     int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "normal case",
			fields: fields{
				Position: &Position{
					NodeId:     1,
					Offset:     0,
					IsReversed: false,
					Name:       "",
				},
				Edits: MakeEditsFromNode(NewNodeFromString(1, "AT")),
				Rank:  1,
			},
			want: "2M",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mapping{
				Position: tt.fields.Position,
				Edits:    tt.fields.Edits,
				Rank:     tt.fields.Rank,
			}
			if got := m.overlapStr(); got != tt.want {
				t.Errorf("overlapStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapping_segmentNameStr(t *testing.T) {
	type fields struct {
		Position *Position
		Edits    []*Edit
		Rank     int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "normal case",
			fields: fields{
				Position: &Position{
					NodeId:     1,
					Offset:     0,
					IsReversed: false,
					Name:       "",
				},
				Edits: MakeEditsFromNode(NewNodeFromString(1, "AT")),
				Rank:  1,
			},
			want: "1+",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mapping{
				Position: tt.fields.Position,
				Edits:    tt.fields.Edits,
				Rank:     tt.fields.Rank,
			}
			if got := m.segmentNameStr(); got != tt.want {
				t.Errorf("segmentNameStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_segmentLine(t *testing.T) {
	type fields struct {
		Id  int64
		Seq *DnaSeq
		Len int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "normal case",
			fields: fields{
				Id:  1,
				Seq: NewDnaSeqFromStr("A"),
				Len: 1,
			},
			want: "S\t1\tA\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Id:  tt.fields.Id,
				Seq: tt.fields.Seq,
				Len: tt.fields.Len,
			}
			if got := n.segmentLine(); got != tt.want {
				t.Errorf("segmentLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPath_pathLine(t *testing.T) {
	type fields struct {
		Name     string
		Mappings []*Mapping
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "normal case",
			fields: fields{
				Name: "spp1",
				Mappings: []*Mapping{
					{
						Position: &Position{1, 0, false, ""},
						Edits:    MakeEditsFromNode(NewNodeFromString(1, "A")),
						Rank:     1,
					},
					{
						Position: &Position{2, 0, false, ""},
						Edits:    MakeEditsFromNode(NewNodeFromString(2, "TC")),
						Rank:     2,
					},
				},
			},
			want: "P\tspp1\t1+,2+\t1M,2M\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Path{
				Name:     tt.fields.Name,
				Mappings: tt.fields.Mappings,
			}
			if got := p.pathLine(); got != tt.want {
				t.Errorf("pathLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerLine(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "normal case",
			want: "H\tVN:Z:1.0\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := headerLine(); got != tt.want {
				t.Errorf("headerLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orientStr(t *testing.T) {
	type args struct {
		isReversed bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal case: positive strand",
			args: args{
				isReversed: false,
			},
			want: "+",
		},
		{
			name: "normal case: negative strand",
			args: args{
				isReversed: true,
			},
			want: "-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orientStr(tt.args.isReversed); got != tt.want {
				t.Errorf("orientStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
