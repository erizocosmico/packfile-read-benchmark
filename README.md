# packfile-read-benchmark

Run benchmarks:

```
make bench
```

### Results

```
go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/erizocosmico/packfile-read-benchmark
BenchmarkDecoder/first-4         	    2000	   1232380 ns/op	   82008 B/op	      60 allocs/op
BenchmarkDecoder/first_100-4     	     100	  13970048 ns/op	 9203436 B/op	    3372 allocs/op
BenchmarkDecoder/first_100_skip_5-4         	     100	  36624806 ns/op	13191426 B/op	    8514 allocs/op
BenchmarkDecoder/last-4                     	    3000	    927435 ns/op	  273406 B/op	     107 allocs/op
BenchmarkRepository/first-4                 	       5	 207959870 ns/op	60300993 B/op	 1238045 allocs/op
BenchmarkRepository/first_100-4             	      10	 257737330 ns/op	74489277 B/op	 1246177 allocs/op
BenchmarkRepository/first_100_skip_5-4      	      10	 251495727 ns/op	78426581 B/op	 1251132 allocs/op
BenchmarkRepository/last-4                  	      10	 153746379 ns/op	60488850 B/op	 1238090 allocs/op
PASS
ok  	github.com/erizocosmico/packfile-read-benchmark	49.624s
```
