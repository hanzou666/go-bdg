package bdg

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	type args struct {
		nodeId int64
		seq    *DnaSeq
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "normal case",
			args: args{nodeId: 1, seq: &DnaSeq{Seq: []byte("ATC")}},
			want: &Node{Id: 1, Seq: &DnaSeq{Seq: []byte("ATC")}, Len: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.nodeId, tt.args.seq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNodeFromString(t *testing.T) {
	type args struct {
		nodeId int64
		seq    string
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "normal case",
			args: args{nodeId: 2, seq: "ATC"},
			want: &Node{
				Id:  2,
				Seq: &DnaSeq{Seq: []byte("ATC")},
				Len: 3,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNodeFromString(tt.args.nodeId, tt.args.seq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNodeFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_fixNodeId(t *testing.T) {
	type fields struct {
		Id  int64
		Seq *DnaSeq
		Len int64
	}
	type args struct {
		baseId int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{
			name: "normal case",
			fields: fields{
				Id:  1,
				Seq: NewDnaSeq("ATC"),
				Len: 3,
			},
			args: args{baseId: 10},
			want: NewNodeFromString(11, "ATC"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Id:  tt.fields.Id,
				Seq: tt.fields.Seq,
				Len: tt.fields.Len,
			}
			if got := n.fixNodeId(tt.args.baseId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fixNodeId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMaxIdNode(t *testing.T) {
	type args struct {
		nodes []*Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "normal case",
			args: args{[]*Node{
				NewNodeFromString(1, "A"),
				NewNodeFromString(100000, "T"),
				NewNodeFromString(3, "C"),
			}},
			want: NewNodeFromString(100000, "T"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxIdNode(tt.args.nodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMaxNodeId() = %v, want %v", got, tt.want)
			}
		})
	}
}
