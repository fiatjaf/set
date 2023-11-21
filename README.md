[![Go Reference](https://pkg.go.dev/badge/github.com/fiatjaf/set.svg)](https://pkg.go.dev/github.com/fiatjaf/set)

Implementation of a simple set using a Go map (`HashSet`) and the same interface using a Go slice (`SliceSet`).

Contrary to what many may think, benchmarks show that the `SliceSet` is faster up to 1000 items (it uses binary search)
and is infinitely faster for returning all the elements inside the set as a slice, of course.

But even with 10000 items it isn't that much slower than `HashSet` for adding elements and for checking for existence
of elements.

```
goos: linux
goarch: amd64
pkg: github.com/fiatjaf/set
cpu: AMD Ryzen 3 3200G with Radeon Vega Graphics
BenchmarkSets/HashSet:10:add()-4         	 3006459	       367.5 ns/op
BenchmarkSets/HashSet:10:has()-4         	21363849	        51.69 ns/op
BenchmarkSets/HashSet:10:slice()-4       	 4780682	       263.5 ns/op
BenchmarkSets/SliceSet:10:add()-4        	 3194888	       356.1 ns/op
BenchmarkSets/SliceSet:10:has()-4        	30230144	        37.17 ns/op
BenchmarkSets/SliceSet:10:slice()-4      	560048186	         2.112 ns/op
BenchmarkSets/HashSet:100:add()-4        	  299668	      4020 ns/op
BenchmarkSets/HashSet:100:has()-4        	22332345	        47.90 ns/op
BenchmarkSets/HashSet:100:slice()-4      	  579247	      2341 ns/op
BenchmarkSets/SliceSet:100:add()-4       	  250821	      5304 ns/op
BenchmarkSets/SliceSet:100:has()-4       	22279682	        46.18 ns/op
BenchmarkSets/SliceSet:100:slice()-4     	563582098	         2.114 ns/op
BenchmarkSets/HashSet:1000:add()-4       	   26728	     44128 ns/op
BenchmarkSets/HashSet:1000:has()-4       	19978395	        58.40 ns/op
BenchmarkSets/HashSet:1000:slice()-4     	   65559	     19742 ns/op
BenchmarkSets/SliceSet:1000:add()-4      	   17730	     64371 ns/op
BenchmarkSets/SliceSet:1000:has()-4      	18923702	        58.03 ns/op
BenchmarkSets/SliceSet:1000:slice()-4    	562180111	         2.104 ns/op
BenchmarkSets/HashSet:10000:add()-4      	    2577	    465568 ns/op
BenchmarkSets/HashSet:10000:has()-4      	20881940	        58.10 ns/op
BenchmarkSets/HashSet:10000:slice()-4    	    7036	    194587 ns/op
BenchmarkSets/SliceSet:10000:add()-4     	    1845	    657179 ns/op
BenchmarkSets/SliceSet:10000:has()-4     	16153288	        73.44 ns/op
BenchmarkSets/SliceSet:10000:slice()-4   	562916802	         2.113 ns/op
```

These sets are not goroutine-safe and they don't implement some classic set operations, like "difference", yet.
