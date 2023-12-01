package set

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
)

type HashSet[V constraints.Ordered] struct {
	internal map[V]bool
}

func NewHashSet[V constraints.Ordered](items ...V) Set[V] {
	internal := make(map[V]bool, len(items))
	for _, item := range items {
		internal[item] = true
	}
	return HashSet[V]{internal: internal}
}

func (s HashSet[V]) Add(items ...V) {
	for _, item := range items {
		s.internal[item] = true
	}
}

func (s HashSet[V]) Has(n V) bool {
	_, ok := s.internal[n]
	return ok
}

func (s HashSet[V]) Remove(items ...V) {
	for _, item := range items {
		delete(s.internal, item)
	}
}

func (s HashSet[V]) Slice() []V {
	return maps.Keys(s.internal)
}

func (s HashSet[V]) Len() int {
	return len(s.internal)
}
