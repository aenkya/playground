package datastructures

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
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
				res1, _ := q.Dequeue().(int)
				res2, _ := q.Dequeue().(int)
				res3, _ := q.Dequeue().(int)
				results = append(results, res1, res2, res3)

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
				res1, _ := q.Peek().(int)
				res2, _ := q.Peek().(int)
				res3, _ := q.Peek().(int)
				results = append(results, res1, res2, res3)

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
