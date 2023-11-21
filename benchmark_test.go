package set

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func BenchmarkSets(b *testing.B) {
	for _, n := range []int{10, 100, 1000, 10000} {
		for _, initSet := range []func(items ...int) Set[int]{NewHashSet[int], NewSliceSet[int]} {
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

func addStuff(s Set[int], n int) {
	for i := 0; i < n; i++ {
		s.Add(i)
	}
}

func checkStuff(s Set[int], n int) {
	_ = s.Has(n * 2)
	_ = s.Has(n / 2)
	_ = s.Has(n * 2 / 3)
	_ = s.Has(n - 1)
	_ = s.Has(n / 4)
}
