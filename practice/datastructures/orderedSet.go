package datastructures

import "sync"

//nolint:govet // ignore struct field order
type OrderedSet struct {
	currentIndex int
	mu           sync.Mutex
	keys         map[any]int
	store        map[any]any
}

func NewOrderedSet() *OrderedSet {
	orderedSet := &OrderedSet{
		keys:  make(map[any]int),
		store: make(map[any]any),
	}

	return orderedSet
}

func (is *OrderedSet) Add(val ...any) {
	is.mu.Lock()
	defer is.mu.Unlock()

	for _, v := range val {
		if _, found := is.keys[v]; found {
			continue
		}

		is.keys[v] = is.currentIndex
		is.store[is.currentIndex] = v

		is.currentIndex++
	}
}

func (is *OrderedSet) Get(key any) (any, bool) {
	is.mu.Lock()
	defer is.mu.Unlock()

	val, ok := is.keys[key]

	if !ok {
		return nil, false
	}

	v, ok := is.store[val]
	if !ok {
		return nil, false
	}

	return v, ok
}

func (is *OrderedSet) Remove(val any) {
	is.mu.Lock()
	defer is.mu.Unlock()

	if index, found := is.store[val]; found {
		delete(is.store, index)
		delete(is.keys, val)
	}
}

func (is *OrderedSet) Reset() {
	is.mu.Lock()
	defer is.mu.Unlock()

	is.currentIndex = 0
	is.keys = make(map[any]int)
	is.store = make(map[any]any)
}

func (is *OrderedSet) Contains(val any) bool {
	if _, found := is.keys[val]; found {
		return true
	}

	return false
}

func (is *OrderedSet) Len() int {
	return len(is.keys)
}

func (is *OrderedSet) Values() []any {
	is.mu.Lock()
	defer is.mu.Unlock()

	vals := make([]any, 0, len(is.keys))

	for k := range is.keys {
		vals = append(vals, k)
	}

	return vals
}

func (is *OrderedSet) Keys() []any {
	is.mu.Lock()
	defer is.mu.Unlock()

	keys := make([]any, 0, len(is.store))

	for i := 0; i < len(is.store); i++ {
		if _, found := is.store[i]; !found {
			continue
		}

		keys = append(keys, i)
	}

	return keys
}
