package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph_AddEdge(t *testing.T) {
	type fields struct {
		graph Graph
	}
	type args struct {
		source      Node
		destination Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Graph
	}{
		{
			name: "Success",
			fields: fields{
				graph: Graph{
					Edges: make(map[Node][]Node),
				},
			},
			args: args{
				source:      "A",
				destination: "B",
			},
			want: Graph{
				Edges: map[Node][]Node{
					"A": {
						"B",
					},
					"B": {
						"A",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				g := &tt.fields.graph

				g.AddEdge(tt.args.source, tt.args.destination)

				assert.Equal(t, tt.want, tt.fields.graph)
			},
		)
	}
}

func TestGraph_IsEulerian(t *testing.T) {
	type fields struct {
		graph Graph
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "IsEulerian",
			fields: fields{
				graph: Graph{
					Edges: map[Node][]Node{
						"A": {
							"C",
							"C",
							"B",
							"B",
						},
						"B": {
							"A",
							"A",
							"D",
						},
						"C": {
							"A",
							"A",
							"D",
						},
						"D": {
							"C",
							"B",
						},
					},
				},
			},
			want: true,
		},
		{
			name: "NotEulerian",
			fields: fields{
				graph: Graph{
					Edges: map[Node][]Node{
						"A": {
							"C",
							"C",
							"B",
							"B",
							"D",
						},
						"B": {
							"A",
							"A",
							"D",
						},
						"C": {
							"A",
							"A",
							"D",
						},
						"D": {
							"C",
							"A",
							"B",
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, tt.fields.graph.IsEulerian())
			},
		)
	}
}
