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
BenchmarkDecoder/first-4         	    2000	   1179847 ns/op	   82089 B/op	      61 allocs/op
BenchmarkDecoder/first_100-4     	     100	  13634571 ns/op	 9208303 B/op	    3472 allocs/op
BenchmarkDecoder/first_100_skip_5-4         	     100	  22796272 ns/op	13195910 B/op	    8614 allocs/op
BenchmarkDecoder/last-4                     	    3000	    787537 ns/op	  273443 B/op	     108 allocs/op
BenchmarkRepository/first-4                 	      10	 145834716 ns/op	60297575 B/op	 1238044 allocs/op
BenchmarkRepository/first_100-4             	      10	 160314738 ns/op	74482661 B/op	 1246176 allocs/op
BenchmarkRepository/first_100_skip_5-4      	      10	 170837738 ns/op	78426575 B/op	 1251132 allocs/op
BenchmarkRepository/last-4                  	      10	 142380387 ns/op	60488852 B/op	 1238090 allocs/op
PASS
ok  	github.com/erizocosmico/packfile-read-benchmark	45.726s
```
