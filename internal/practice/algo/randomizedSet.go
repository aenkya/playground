package algo

import (
	"crypto/rand"
	"math/big"
)

type RandomizedSet struct {
	values map[int]int
	keys   []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		keys:   []int{},
		values: make(map[int]int),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.values[val]; ok {
		return false
	}

	rs.keys = append(rs.keys, val)
	rs.values[val] = len(rs.keys) - 1

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	i, exists := rs.values[val]

	if !exists {
		return false
	}

	last := rs.keys[len(rs.keys)-1]
	rs.keys[i] = last
	rs.values[last] = i

	rs.keys = rs.keys[:len(rs.keys)-1]

	delete(rs.values, val)

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	rn, _ := rand.Int(rand.Reader, big.NewInt(int64(len(rs.keys))))
	return rs.keys[rn.Int64()]
}
