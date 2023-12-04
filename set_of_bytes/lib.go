package set_of_bytes

type Set interface {
	Add(item []byte)
	Remove(item []byte)
	Has(item []byte) bool
	Slice() [][]byte
	Len() int
}

var (
	_ Set = (*SliceSet)(nil)
	_ Set = (*HashSet)(nil)
)
