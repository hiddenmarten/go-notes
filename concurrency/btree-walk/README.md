# Helpful links:

- https://go.dev/tour/concurrency/7
- https://www.practical-go-lessons.com/chap-34-benchmarks

Benchmark run:
```shell
go test -bench=. -benchmem -timeout 30m > benchmark_results.txt
```

Plots generation:
```shell
go run plots.go
```

### Description

**Objective**: To compare the efficiency of a worker pool versus goroutines on demand in handling varying loads per operation.

**Setup**:
- Implement two systems: one using a worker pool and the other using goroutines on demand.
- Use both systems to perform a series of operations that involve deploying binary trees.

**Parameters**:
- Measure resource utilization, including CPU and memory usage.
- Record overhead and useful load to evaluate efficiency.

**Observations**:
- **Small Load Per Operation**: Worker pools show better resource utilization than goroutines on demand. Overhead in the latter can exceed the useful load, making it less efficient.
- **Large Load Per Operation**: Goroutines on demand are more efficient, showing better performance and resource utilization.

**Conclusion**:
- Worker pools are beneficial in scenarios with a low and medium load per operation, requiring a buffered task channel for optimal performance.
- Goroutines on demand are more efficient in high-load scenarios per operation.
- Implementing a worker pool is more complex and requires scaling with increased load, potentially leading to more complex code. Automating scaling can help but adds to the complexity.
