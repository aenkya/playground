package datastructures

import (
	"testing"

	"enkya.org/playground/internal/utils"
)

func TestSinglyLinkedList(t *testing.T) {
	tests := []struct {
		name string
		list *SinglyLinkedList
		want []int
	}{
		{
			name: "empty list",
			list: NewSinglyLinkedList(),
			want: []int{},
		},
		{
			name: "append",
			list: func() *SinglyLinkedList {
				list := NewSinglyLinkedList()
				list.Append(1)
				list.Append(2)
				list.Append(3)
				return list
			}(),
			want: []int{1, 2, 3},
		},
		{
			name: "prepend",
			list: func() *SinglyLinkedList {
				list := NewSinglyLinkedList()
				list.Prepend(1)
				list.Prepend(2)
				list.Prepend(3)
				return list
			}(),
			want: []int{3, 2, 1},
		},
		{
			name: "insert at",
			list: func() *SinglyLinkedList {
				list := NewSinglyLinkedList()
				list.Append(1)
				list.Append(3)
				list.InsertAt(1, 2)
				return list
			}(),
			want: []int{1, 2, 3},
		},
		{
			name: "remove at",
			list: func() *SinglyLinkedList {
				list := NewSinglyLinkedList()
				list.Append(1)
				list.Append(2)
				list.Append(3)
				list.RemoveAt(1)
				return list
			}(),
			want: []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.list.ToArray()
			got := make([]int, len(g))

			for i, e := range g {
				got[i], _ = e.(int)
			}

			if !utils.CompareSlice(tt.want, got) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
