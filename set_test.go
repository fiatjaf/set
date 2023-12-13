package set

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSets(t *testing.T) {
	for _, initSet := range []func(items ...string) Set[string]{NewHashSet[string], NewSliceSet[string]} {
		s := initSet()
		typ := strings.Split(
			strings.Split(
				reflect.TypeOf(s).String(),
				".")[1],
			"[",
		)[0]

		t.Run(fmt.Sprintf("%s:add/len/remove/has/slice", typ), func(t *testing.T) {
			s.Add("a")
			s.Add("a", "b", "c")
			s.Add("b", "c")
			if s.Len() != 3 {
				t.Fatalf("Len() is wrong")
			}
			s.Add("d")
			s.Add("h")
			s.Add("e", "f", "g", "h")
			if s.Len() != 8 {
				t.Fatalf("Len() is wrong")
			}

			for _, l := range []string{"a", "b", "c", "d", "e", "f", "g", "h"} {
				if !s.Has(l) {
					t.Fatalf("'%s' should be in the set", l)
				}
			}
			for _, l := range []string{"psow", "xucj"} {
				if s.Has(l) {
					t.Fatalf("'%s' should not be in the set", l)
				}
			}

			s.Add("xucj", "psow", "ziiz", "ploc")
			s.Remove("e", "psow", "a", "h")
			s.Remove("ziiz")

			for _, l := range []string{"b", "c", "d", "f", "g", "xucj", "ploc"} {
				if !s.Has(l) {
					t.Fatalf("'%s' should be in the set", l)
				}
			}
			for _, l := range []string{"psow", "a", "e", "h", "ziiz"} {
				if s.Has(l) {
					t.Fatalf("'%s' should not be in the set", l)
				}
			}

			if len(s.Slice()) != s.Len() {
				t.Fatalf("slice has a different length")
			}

			if s.Len() != NewSliceSet(s.Slice()...).Len() {
				t.Fatalf("dup is different than original")
			}
		})
	}
}
