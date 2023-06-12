package utils

import "sync"

//nolint:govet // ignore struct field order
type OrderedSet struct {
	mu    sync.Mutex
	keys  []any
	store map[any]any
}

func NewOrderedSet() *OrderedSet {
	orderedSet := &OrderedSet{
		keys:  make([]any, 0),
		store: make(map[any]any),
	}

	return orderedSet
}

func (is *OrderedSet) Add(key, val any) {
	is.mu.Lock()
	defer is.mu.Unlock()

	if _, ok := is.store[key]; !ok {
		is.keys = append(is.keys, key)
	}

	is.store[key] = val
}

func (is *OrderedSet) Remove(key any) {
	is.mu.Lock()
	defer is.mu.Unlock()

	if _, found := is.store[key]; !found {
		return
	}

	delete(is.store, key)

	for k, v := range is.keys {
		if v == key {
			is.keys = append(is.keys[:k], is.keys[k+1:]...)
		}
	}
}

func (is *OrderedSet) Contains(val any) bool {
	for _, v := range is.store {
		if v == val {
			return true
		}
	}

	return false
}
