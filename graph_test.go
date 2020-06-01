package bdg

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	g1 := NewGraph(
		[]*Node{NewNodeFromString(1, "A"), NewNodeFromString(2, "T")},
		[]*Edge{{1, 2, false, false}},
		[]*Path{{
			Name: "spp1",
			Mappings: []*Mapping{
				{&Position{1, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(1, "A")), 1},
				{&Position{2, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(2, "T")), 2},
			}},
		},
	)
	g2 := NewGraph(
		[]*Node{NewNodeFromString(1, "A"), NewNodeFromString(2, "T")},
		[]*Edge{{1, 2, false, false}},
		[]*Path{{
			Name: "spp2",
			Mappings: []*Mapping{
				{&Position{1, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(1, "A")), 1},
				{&Position{2, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(2, "T")), 2},
			}},
		},
	)

	type args struct {
		graphs []*Graph
	}
	tests := []struct {
		name string
		args args
		want *Graph
	}{
		{
			name: "normal case",
			args: args{
				graphs: []*Graph{g1, g2},
			},
			want: NewGraph(
				[]*Node{
					NewNodeFromString(1, "A"),
					NewNodeFromString(2, "T"),
					NewNodeFromString(3, "A"),
					NewNodeFromString(4, "T"),
				}, []*Edge{
					{1, 2, false, false},
					{3, 4, false, false},
				}, []*Path{
					{
						Name: "spp1",
						Mappings: []*Mapping{
							{&Position{1, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(1, "A")), 1},
							{&Position{2, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(2, "T")), 2},
						},
					}, {
						Name: "spp2",
						Mappings: []*Mapping{
							{&Position{3, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(3, "A")), 1},
							{&Position{4, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(4, "T")), 2},
						},
					},
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Combine(tt.args.graphs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Combine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_GetPath(t *testing.T) {
	g := NewGraph(
		[]*Node{NewNodeFromString(1, "A"), NewNodeFromString(2, "T")},
		[]*Edge{{1, 2, false, false}},
		[]*Path{{
			Name: "spp1",
			Mappings: []*Mapping{
				{&Position{1, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(1, "A")), 1},
				{&Position{2, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(2, "T")), 2},
			}},
		},
	)
	type fields struct {
		Nodes     []*Node
		Edges     []*Edge
		Paths     []*Path
		nodeIndex map[int64]*Node
		pathIndex map[string]*Path
	}
	type args struct {
		pathName string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantPath *Path
		wantOk   bool
	}{
		{
			name: "normal case",
			fields: fields{
				Nodes:     g.Nodes,
				Edges:     g.Edges,
				Paths:     g.Paths,
				nodeIndex: g.nodeIndex,
				pathIndex: g.pathIndex,
			},
			args: args{pathName: "spp1"},
			wantPath: &Path{
				Name: "spp1",
				Mappings: []*Mapping{
					{&Position{1, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(1, "A")), 1},
					{&Position{2, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(2, "T")), 2},
				},
			},
			wantOk: true,
		},
		{
			name: "not found case",
			fields: fields{
				Nodes:     g.Nodes,
				Edges:     g.Edges,
				Paths:     g.Paths,
				nodeIndex: g.nodeIndex,
				pathIndex: g.pathIndex,
			},
			args:     args{pathName: "spp2"},
			wantPath: nil,
			wantOk:   false,
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
			gotPath, gotOk := g.GetPath(tt.args.pathName)
			if !reflect.DeepEqual(gotPath, tt.wantPath) {
				t.Errorf("GetPath() gotPath = %v, want %v", gotPath, tt.wantPath)
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetPath() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestNewGraph(t *testing.T) {
	type args struct {
		nodes []*Node
		edges []*Edge
		paths []*Path
	}
	tests := []struct {
		name string
		args args
		want *Graph
	}{
		// TODO: Add test cases.
		{
			name: "normal case",
			args: args{
				nodes: []*Node{NewNodeFromString(1, "A"), NewNodeFromString(2, "T")},
				edges: []*Edge{{1, 2, false, false}},
				paths: []*Path{{
					Name: "spp1",
					Mappings: []*Mapping{
						{&Position{1, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(1, "A")), 1},
						{&Position{2, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(2, "T")), 2},
					}},
				},
			},
			want: &Graph{
				Nodes: []*Node{NewNodeFromString(1, "A"), NewNodeFromString(2, "T")},
				Edges: []*Edge{{1, 2, false, false}},
				Paths: []*Path{{
					Name: "spp1",
					Mappings: []*Mapping{
						{&Position{1, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(1, "A")), 1},
						{&Position{2, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(2, "T")), 2},
					}},
				},
				nodeIndex: map[int64]*Node{
					1: NewNodeFromString(1, "A"),
					2: NewNodeFromString(2, "T"),
				},
				pathIndex: map[string]*Path{
					"spp1": &Path{
						Name: "spp1",
						Mappings: []*Mapping{
							{&Position{1, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(1, "A")), 1},
							{&Position{2, 0, false, ""}, MakeEditsFromNode(NewNodeFromString(2, "T")), 2},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGraph(tt.args.nodes, tt.args.edges, tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
