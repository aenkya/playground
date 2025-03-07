package datastructures

import (
	"reflect"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tests := []struct {
		name      string
		operation func() []any
		want      []any
	}{
		{
			name: "insert",
			operation: func() []any {
				tree := NewBinaryTree()

				tree.Insert(5)
				tree.Insert(3)
				tree.Insert(7)
				tree.Insert(2)
				tree.Insert(4)
				tree.Insert(6)
				tree.Insert(8)

				return tree.DepthFirstTraversal(InOrder)
			},
			want: []any{2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "search",
			operation: func() []any {
				tree := NewBinaryTree()

				tree.Insert(5)
				tree.Insert(3)
				tree.Insert(7)
				tree.Insert(2)
				tree.Insert(4)
				tree.Insert(6)
				tree.Insert(8)

				t := tree.Search(7)
				return []any{t.Value}
			},
			want: []any{7},
		},
		{
			name: "delete",
			operation: func() []any {
				tree := NewBinaryTree()

				tree.Insert(5)
				tree.Insert(3)
				tree.Insert(7)
				tree.Insert(2)
				tree.Insert(4)
				tree.Insert(6)
				tree.Insert(8)

				tree.Delete(7)

				return tree.DepthFirstTraversal(InOrder)
			},
			want: []any{2, 3, 4, 5, 6, 8},
		},
		{
			name: "depth first traversal",
			operation: func() []any {
				tree := NewBinaryTree()

				tree.Insert(5)
				tree.Insert(3)
				tree.Insert(7)
				tree.Insert(2)
				tree.Insert(4)
				tree.Insert(6)
				tree.Insert(8)

				return tree.DepthFirstTraversal(InOrder)
			},
			want: []any{2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "breadth first traversal",
			operation: func() []any {
				tree := NewBinaryTree()

				tree.Insert(5)
				tree.Insert(3)
				tree.Insert(7)
				tree.Insert(2)
				tree.Insert(4)
				tree.Insert(6)
				tree.Insert(8)

				return tree.BreadthFirstTraversal()
			},
			want: []any{5, 3, 7, 2, 4, 6, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := tt.operation()

			if !reflect.DeepEqual(results, tt.want) {
				t.Errorf("got %v, want %v", results, tt.want)
			}
		})
	}
}
