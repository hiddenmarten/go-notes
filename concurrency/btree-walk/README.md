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
- NoPool:
    - With a small load per operation, the overhead of goroutines on demand can exceed the useful load.
    - The number of `allocs/op` increases with the depth of the investigated tree.
    - The number of `B/op` increases with the depth of the investigated tree.
- Pool:
    - The number of `allocs/op` barely changes in this experiment for the worker pool, despite the changed payload.
    - The number of `B/op` barely changes in this experiment, except at the point when tasks accumulate in the buffered channel.

**Conclusions**:
- A worker pool utilizes resources more stably and predictably than goroutines on demand.
- A worker pool requires a buffered task channel for effective operation.
- A worker pool is more complex to implement.
- A worker pool needs to be scaled according to the increasing load, which can be automated, but this leads to even more complex code.
- A no-pool solution looks better with "unlimited" resources.
- With a large load per operation, goroutines on demand appear more efficient, at least in this experiment involving the deployment of binary trees where the number of active tasks increases progressively.
