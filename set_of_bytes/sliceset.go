package set_of_bytes

import (
	"golang.org/x/exp/slices"
)

type SliceSet struct {
	items [][]byte
}

func NewSliceSet() Set {
	return &SliceSet{items: make([][]byte, 0, 12)}
}

func (s *SliceSet) Add(item []byte) {
	idx, exists := slices.BinarySearchFunc(s.items, item, func(a1, a2 []byte) int {
		return slices.Compare(a1, a2)
	})
	if exists {
		return
	}
	s.items = append(s.items, item) // bogus append just to increase the capacity
	copy(s.items[idx+1:], s.items[idx:])
	s.items[idx] = item
}

func (s *SliceSet) Has(item []byte) bool {
	_, exists := slices.BinarySearchFunc(s.items, item, func(a1 []byte, a2 []byte) int {
		return slices.Compare(a1, a2)
	})
	return exists
}

func (s *SliceSet) Remove(item []byte) {
	idx, exists := slices.BinarySearchFunc(s.items, item, func(a1, a2 []byte) int {
		return slices.Compare(a1, a2)
	})
	if !exists {
		return
	}
	copy(s.items[idx:], s.items[idx-1:])
	s.items = s.items[0 : len(s.items)-1]
}

func (s *SliceSet) Slice() [][]byte {
	return s.items
}

func (s *SliceSet) Len() int {
	return len(s.items)
}
