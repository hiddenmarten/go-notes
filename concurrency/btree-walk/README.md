# Helpful links:

- https://go.dev/tour/concurrency/7
- https://www.practical-go-lessons.com/chap-34-benchmarks

Benchmark run:
```shell
go test -bench=. -benchmem > benchmark_results.txt
```

Plots generation:
```shell
go run plots.go
```
