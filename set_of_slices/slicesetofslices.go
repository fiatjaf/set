package set_of_slices

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type SetOfSlices[A constraints.Ordered] struct {
	items [][]A
}

func NewSetOfSlices[A constraints.Ordered](items ...[]A) SetOfSlices[A] {
	slices.SortFunc(items, func(a, b []A) int {
		return slices.Compare(a, b)
	})
	return SetOfSlices[A]{items: items}
}

func (s *SetOfSlices[A]) Add(items ...[]A) {
	for _, a := range items {
		idx, exists := slices.BinarySearchFunc(s.items, a, func(a1, a2 []A) int {
			return slices.Compare(a1, a2)
		})
		if exists {
			return
		}
		s.items = append(s.items, a) // bogus append just to increase the capacity
		copy(s.items[idx+1:], s.items[idx:])
		s.items[idx] = a
	}
}

func (s *SetOfSlices[A]) Has(item []A) bool {
	_, exists := slices.BinarySearchFunc(s.items, item, func(a1 []A, a2 []A) int {
		return slices.Compare(a1, a2)
	})
	return exists
}

func (s *SetOfSlices[A]) Remove(items ...[]A) {
	for _, a := range items {
		idx, exists := slices.BinarySearchFunc(s.items, a, func(a1, a2 []A) int {
			return slices.Compare(a1, a2)
		})
		if !exists {
			return
		}
		copy(s.items[idx:], s.items[idx-1:])
		s.items = s.items[0 : len(s.items)-1]
	}
}

func (s *SetOfSlices[A]) Slice() [][]A {
	return s.items
}

func (s *SetOfSlices[A]) Len() int {
	return len(s.items)
}
