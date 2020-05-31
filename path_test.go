package bdg

import (
	"reflect"
	"testing"
)

func TestNewEditOfNode(t *testing.T) {
	type args struct {
		n *Node
	}
	tests := []struct {
		name string
		args args
		want *Edit
	}{
		{
			name: "normal case",
			args: args{
				n: NewNodeFromString(1, "ATC"),
			},
			want: &Edit{
				FromLength: 3,
				ToLength:   3,
				Sequence:   NewDnaSeqFromStr("ATC"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEditOfNode(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEditOfNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_fixNodeId(t *testing.T) {
	type fields struct {
		NodeId     int64
		Offset     int64
		IsReversed bool
		Name       string
	}
	type args struct {
		baseId int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Position
	}{
		{
			name:   "normal case",
			fields: fields{NodeId: 1, Offset: 0, IsReversed: false, Name: "pos1"},
			args:   args{10},
			want: &Position{
				NodeId:     11,
				Offset:     0,
				IsReversed: false,
				Name:       "pos1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				NodeId:     tt.fields.NodeId,
				Offset:     tt.fields.Offset,
				IsReversed: tt.fields.IsReversed,
				Name:       tt.fields.Name,
			}
			if got := p.fixNodeId(tt.args.baseId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fixNodeId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapping_fixNodeId(t *testing.T) {
	type fields struct {
		Position *Position
		Edits    []*Edit
		Rank     int64
	}
	type args struct {
		baseId int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Mapping
	}{
		{
			name: "normal case",
			fields: fields{
				Position: &Position{
					NodeId:     1,
					Offset:     0,
					IsReversed: false,
					Name:       "pos1",
				},
				Edits: []*Edit{NewEditOfNode(NewNodeFromString(1, "ATG"))},
				Rank:  1,
			},
			args: args{baseId: 10},
			want: &Mapping{
				Position: &Position{
					NodeId:     11,
					Offset:     0,
					IsReversed: false,
					Name:       "pos1",
				},
				Edits: []*Edit{NewEditOfNode(NewNodeFromString(11, "ATG"))},
				Rank:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mapping{
				Position: tt.fields.Position,
				Edits:    tt.fields.Edits,
				Rank:     tt.fields.Rank,
			}
			if got := m.fixNodeId(tt.args.baseId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fixNodeId() = %v, want %v", got, tt.want)
			}
		})
	}
}
