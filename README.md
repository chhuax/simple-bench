# simple-bench
simple benchmark code of different kv databases in Go.

## Test database:

* [bbolt](https://github.com/etcd-io/bbolt)
* [goleveldb](https://github.com/syndtr/goleveldb)
* [nutsdb](https://github.com/nutsdb/nutsdb)
* [rosedb](https://github.com/flower-corp/rosedb)

## Options:

* No sync
* Value size: 512 bytes

## Results

`go test -v -bench=.` 

```
goos: linux
goarch: amd64
pkg: rosedb-bench
cpu: Intel(R) Xeon(R) Platinum 8255C CPU @ 2.50GHz
BenchmarkPutValue_BoltDB
BenchmarkPutValue_BoltDB-2      	   22291	     52287 ns/op	   20735 B/op	      75 allocs/op
BenchmarkGetValue_BoltDB
BenchmarkGetValue_BoltDB-2      	  423628	      2465 ns/op	     599 B/op	      10 allocs/op
BenchmarkPutValue_GoLevelDB
BenchmarkPutValue_GoLevelDB-2   	   69589	     19481 ns/op	    2340 B/op	      12 allocs/op
BenchmarkGetValue_GoLevelDB
BenchmarkGetValue_GoLevelDB-2   	  265510	      4531 ns/op	    1274 B/op	      15 allocs/op
BenchmarkPutValue_NutsDB
BenchmarkPutValue_NutsDB-2      	   72456	     18502 ns/op	    3503 B/op	      22 allocs/op
BenchmarkGetValue_NutsDB
BenchmarkGetValue_NutsDB-2      	  157114	      6746 ns/op	     816 B/op	      12 allocs/op
BenchmarkPutValue_RoseDB
BenchmarkPutValue_RoseDB-2      	   65773	     18870 ns/op	    3387 B/op	      13 allocs/op
BenchmarkGetValue_RoseDB
BenchmarkGetValue_RoseDB-2      	  371708	      3496 ns/op	     743 B/op	       6 allocs/op
```

`go test -v -bench=. -benchtime=1000000x`

```
goos: linux
goarch: amd64
pkg: rosedb-bench
cpu: Intel(R) Xeon(R) Platinum 8255C CPU @ 2.50GHz
BenchmarkPutValue_BoltDB
BenchmarkPutValue_BoltDB-2      	 1000000	     61584 ns/op	   26307 B/op	      77 allocs/op
BenchmarkGetValue_BoltDB
BenchmarkGetValue_BoltDB-2      	 1000000	      2398 ns/op	     600 B/op	      10 allocs/op
BenchmarkPutValue_GoLevelDB
BenchmarkPutValue_GoLevelDB-2   	 1000000	     18769 ns/op	    2255 B/op	      12 allocs/op
BenchmarkGetValue_GoLevelDB
BenchmarkGetValue_GoLevelDB-2   	 1000000	      4529 ns/op	    1284 B/op	      15 allocs/op
BenchmarkPutValue_NutsDB
BenchmarkPutValue_NutsDB-2      	 1000000	     20205 ns/op	    3553 B/op	      22 allocs/op
BenchmarkGetValue_NutsDB
BenchmarkGetValue_NutsDB-2      	 1000000	     13378 ns/op	    1471 B/op	      19 allocs/op
BenchmarkPutValue_RoseDB
BenchmarkPutValue_RoseDB-2      	 1000000	     18621 ns/op	    3445 B/op	      15 allocs/op
BenchmarkGetValue_RoseDB
BenchmarkGetValue_RoseDB-2      	 1000000	      3549 ns/op	     743 B/op	       6 allocs/op
```

