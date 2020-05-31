package bdg

import (
	"reflect"
	"testing"
)

func TestEdge_fixNodeId(t *testing.T) {
	type fields struct {
		FromId    int64
		ToId      int64
		FromStart bool
		ToEnd     bool
	}
	type args struct {
		baseId int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Edge
	}{
		{
			name: "normal case",
			fields: fields{
				FromId:    1,
				ToId:      2,
				FromStart: false,
				ToEnd:     false,
			},
			args: args{baseId: 10},
			want: &Edge{
				FromId:    11,
				ToId:      12,
				FromStart: false,
				ToEnd:     false,
			},
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
			if got := e.fixNodeId(tt.args.baseId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fixNodeId() = %v, want %v", got, tt.want)
			}
		})
	}
}
