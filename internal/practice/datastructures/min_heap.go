package datastructures

import (
	"errors"
	"fmt"
)

type MinHeap struct {
	data        []int
	positionMap map[int]int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data:        []int{},
		positionMap: make(map[int]int),
	}
}

func (h *MinHeap) Push(val int) {
	h.data = append(h.data, val)
	index := len(h.data) - 1
	h.positionMap[val] = index
	h.heapifyUp(index)
}

func (h *MinHeap) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[i] < h.data[parent] {
			h.swap(i, parent)
			i = parent
		} else {
			break
		}
	}
}

func (h *MinHeap) swap(i, j int) {
	h.positionMap[h.data[i]], h.positionMap[h.data[j]] = j, i
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinHeap) Pop() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("heap is empty")
	}

	minn := h.data[0]
	last := h.data[len(h.data)-1]
	h.data[0] = last
	h.positionMap[last] = 0
	delete(h.positionMap, minn)
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown(0)

	return minn, nil
}

func (h *MinHeap) heapifyDown(i int) {
	n := len(h.data)

	for {
		smallest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && h.data[left] < h.data[smallest] {
			smallest = left
		}

		if right < n && h.data[right] < h.data[smallest] {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.swap(i, smallest)
		i = smallest
	}
}

func (h *MinHeap) DecreaseKey(oldVal, newVal int) error {
	index, exists := h.positionMap[oldVal]
	if !exists {
		return errors.New("value not found")
	}

	if newVal > oldVal {
		return errors.New("new value is greater than current value")
	}

	h.data[index] = newVal
	delete(h.positionMap, oldVal)
	h.positionMap[newVal] = index
	h.heapifyUp(index)

	return nil
}

func (h *MinHeap) Print() {
	fmt.Println("Heap: ", h.data)
}
