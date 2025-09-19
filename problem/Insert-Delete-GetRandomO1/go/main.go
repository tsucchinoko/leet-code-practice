package main

import (
	"math/rand"
	"time"
)

// RandomizedSet implements insert, remove, getRandom in average O(1).
type RandomizedSet struct {
	values []int
	idxMap map[int]int
	rnd    *rand.Rand
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		values: make([]int, 0),
		idxMap: make(map[int]int),
		rnd:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.idxMap[val]; ok {
		return false
	}
	rs.idxMap[val] = len(rs.values)
	rs.values = append(rs.values, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	idx, ok := rs.idxMap[val]
	if !ok {
		return false
	}
	lastIdx := len(rs.values) - 1
	lastVal := rs.values[lastIdx]

	// Move last element to idx (unless idx is lastIdx)
	rs.values[idx] = lastVal
	rs.idxMap[lastVal] = idx

	// Pop last
	rs.values = rs.values[:lastIdx]
	delete(rs.idxMap, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	n := len(rs.values)
	// Assumes n > 0
	i := rs.rnd.Intn(n)
	return rs.values[i]
}
