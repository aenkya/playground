package algo

import (
	"container/heap"
)

type MaxIntHeap []int

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxIntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxIntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MaxIntHeap{}
	heap.Init(h)

	for _, v := range nums {
		heap.Push(h, v)
	}

	v := 0
	for range k {
		v = heap.Pop(h).(int)
	}

	return v
}
