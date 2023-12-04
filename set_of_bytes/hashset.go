package set_of_bytes

type HashSet struct {
	internal map[string]bool
}

func NewHashSet() Set {
	internal := make(map[string]bool)
	return HashSet{internal: internal}
}

func (s HashSet) Add(item []byte) {
	s.internal[string(item)] = true
}

func (s HashSet) Has(item []byte) bool {
	_, ok := s.internal[string(item)]
	return ok
}

func (s HashSet) Remove(item []byte) {
	delete(s.internal, string(item))
}

func (s HashSet) Slice() [][]byte {
	items := make([][]byte, len(s.internal))
	i := 0
	for sitem := range s.internal {
		items[i] = []byte(sitem)
		i++
	}
	return items
}

func (s HashSet) Len() int {
	return len(s.internal)
}
