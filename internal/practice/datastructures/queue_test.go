package datastructures

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	//nolint:govet // ignore struct optimisation
	tests := []struct {
		name      string
		operation func() any
		want      interface{}
	}{
		{
			name: "enqueue and dequeue",
			operation: func() any {
				q := NewQueue()

				q.Enqueue(1)
				q.Enqueue(2)
				q.Enqueue(3)

				results := []int{}
				results = append(results, q.Dequeue().(int), q.Dequeue().(int), q.Dequeue().(int))

				return results
			},
			want: []int{1, 2, 3},
		},
		{
			name: "peek",
			operation: func() any {
				q := NewQueue()

				q.Enqueue(1)
				q.Enqueue(2)
				q.Enqueue(3)

				results := []int{}
				results = append(results, q.Peek().(int), q.Peek().(int), q.Peek().(int))

				return results
			},
			want: []int{1, 1, 1},
		},
		{
			name: "len",
			operation: func() any {
				q := NewQueue()

				q.Enqueue(1)
				q.Enqueue(2)
				q.Enqueue(3)

				results := []int{}
				results = append(results, q.Len())
				_ = q.Dequeue()
				results = append(results, q.Len())

				return results
			},
			want: []int{3, 2},
		},
		{
			name: "is empty",
			operation: func() any {
				q := NewQueue()

				q.Enqueue(1)
				q.Enqueue(2)

				results := []bool{}
				results = append(results, q.IsEmpty())
				_ = q.Dequeue()
				results = append(results, q.IsEmpty())
				_ = q.Dequeue()
				results = append(results, q.IsEmpty())

				return results
			},
			want: []bool{false, false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.operation()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
