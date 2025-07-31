package datastructures

import "errors"

type MaxHeap struct {
	data        []int
	positionMap map[int]int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data:        []int{},
		positionMap: make(map[int]int),
	}
}

func (h *MaxHeap) Push(val int) {
	h.data = append(h.data, val)
	index := len(h.data) - 1
	h.positionMap[val] = index
	h.heapifyUp(index)
}

func (h *MaxHeap) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[i] > h.data[parent] {
			h.swap(i, parent)
			i = parent
		} else {
			break
		}
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.positionMap[h.data[i]], h.positionMap[h.data[j]] = j, i
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MaxHeap) Max() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("heap is empty")
	}

	return h.data[0], nil
}

func (h *MaxHeap) Pop() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("heap is empty")
	}

	maxn := h.data[0]
	last := h.data[len(h.data)-1]

	h.data[0] = last
	h.positionMap[last] = 0
	delete(h.positionMap, maxn)

	h.data = h.data[:len(h.data)-1]
	h.heapifyDown(0)

	return maxn, nil
}

func (h *MaxHeap) heapifyDown(i int) {
	n := len(h.data)

	for {
		largest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && h.data[left] > h.data[largest] {
			largest = left
		}

		if right < n && h.data[right] > h.data[largest] {
			largest = right
		}

		if largest == i {
			break
		}

		h.swap(i, largest)
		i = largest
	}
}

func (h *MaxHeap) IncreaseKey(oldVal, newVal int) error {
	index, exists := h.positionMap[oldVal]
	if !exists {
		return errors.New("value not found in heap")
	}

	if newVal < oldVal {
		return errors.New("new value is less than current value")
	}

	h.data[index] = newVal
	delete(h.positionMap, oldVal)
	h.positionMap[newVal] = index
	h.heapifyUp(index)
	return nil
}
