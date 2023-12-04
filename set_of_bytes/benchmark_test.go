package set_of_bytes

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

var rng *rand.Rand

func BenchmarkSets(b *testing.B) {
	rng = rand.New(rand.NewSource(1))
	for _, n := range []int{10, 100, 1000, 10000} {
		for _, initSet := range []func() Set{NewHashSet, NewSliceSet} {
			s := initSet()
			typ := strings.Split(
				strings.Split(
					reflect.TypeOf(s).String(),
					".")[1],
				"[",
			)[0]

			b.Run(fmt.Sprintf("%s:%d:add()", typ, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					addStuff(s, n)
				}
			})
			b.Run(fmt.Sprintf("%s:%d:has()", typ, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					checkStuff(s, n)
				}
			})
			b.Run(fmt.Sprintf("%s:%d:slice()", typ, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = s.Slice()
				}
			})
		}
	}
}

func addStuff(s Set, n int) {
	for i := 0; i < n; i++ {
		arr := make([]byte, 32)
		rng.Read(arr)
		s.Add(arr)
	}
}

func checkStuff(s Set, n int) {
	for i := 0; i < n; i++ {
		arr := make([]byte, 32)
		rng.Read(arr)
		s.Add(arr)
	}
}
