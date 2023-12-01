package set

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type SliceSet[A constraints.Ordered] struct {
	items []A
}

func NewSliceSet[A constraints.Ordered](items ...A) Set[A] {
	slices.Sort(items)
	return &SliceSet[A]{items: items}
}

func (s *SliceSet[A]) Add(items ...A) {
	for _, a := range items {
		idx, exists := slices.BinarySearch(s.items, a)
		if exists {
			return
		}
		s.items = append(s.items, a) // bogus append just to increase the capacity
		copy(s.items[idx+1:], s.items[idx:])
		s.items[idx] = a
	}
}

func (s *SliceSet[A]) Has(item A) bool {
	_, exists := slices.BinarySearch(s.items, item)
	return exists
}

func (s *SliceSet[A]) Remove(items ...A) {
	for _, a := range items {
		idx, exists := slices.BinarySearch(s.items, a)
		if !exists {
			return
		}
		copy(s.items[idx:], s.items[idx-1:])
		s.items = s.items[0 : len(s.items)-1]
	}
}

func (s *SliceSet[A]) Slice() []A {
	return s.items
}

func (s *SliceSet[A]) Len() int {
	return len(s.items)
}
