package datastructures

import (
	"fmt"
	"testing"
)

func TestOrderedSet(t *testing.T) {
	tests := []struct {
		name string
		set  *OrderedSet
		want []int
	}{
		{
			name: "empty set",
			set:  NewOrderedSet(),
			want: []int{},
		},
		{
			name: "add",
			set: func() *OrderedSet {
				set := NewOrderedSet()
				set.Add(1)
				set.Add(2)
				set.Add(3)
				return set
			}(),
			want: []int{1, 2, 3},
		},
		{
			name: "get",
			set: func() *OrderedSet {
				set := NewOrderedSet()
				set.Add(1)
				set.Add(2)
				set.Add(3)
				return set
			}(),
			want: []int{1, 2, 3},
		},
		{
			name: "remove",
			set: func() *OrderedSet {
				set := NewOrderedSet()
				set.Add(1)
				set.Add(2)
				set.Add(3)
				set.Remove(2)
				return set
			}(),
			want: []int{1, 3},
		},
		{
			name: "reset",
			set: func() *OrderedSet {
				set := NewOrderedSet()
				set.Add(1)
				set.Add(2)
				set.Add(3)
				set.Reset()
				return set
			}(),
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for _, v := range tt.want {
				fmt.Println(tt.set.Get(v))
				if _, found := tt.set.Get(v); !found {
					t.Errorf("expected to find %v in set", v)
				}
			}
		})
	}
}
