package main

import (
	"sync"
	"time"
)

// Task represents a tree node to be processed.
type Task struct {
	t  *Tree
	ch chan int
	wg *sync.WaitGroup
}

// WorkerPool manages a pool of goroutines to process tree nodes.
type WorkerPool struct {
	tasks       chan Task
	workerCount int
	wg          sync.WaitGroup
}

// NewWorkerPool initializes a new worker pool.
func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		// This channel MUST be buffered to make goroutines work efficient.
		// One more thing is that with no buffered channel and low number of workers you'll face deadlock.
		tasks:       make(chan Task, tasksChanBufferSize),
		workerCount: workerCount,
	}
}

// Start initializes the goroutines and begins processing tasks.
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
}

// Stop waits for all goroutines to finish processing tasks.
func (wp *WorkerPool) Stop() {
	close(wp.tasks)
	wp.wg.Wait()
}

// worker is a goroutine that processes tasks from the task channel.
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for task := range wp.tasks {
		time.Sleep(pseudoPayload)
		task.ch <- task.t.Value
		if task.t.Left != nil {
			task.wg.Add(1)
			wp.tasks <- Task{task.t.Left, task.ch, task.wg}
		}
		if task.t.Right != nil {
			task.wg.Add(1)
			wp.tasks <- Task{task.t.Right, task.ch, task.wg}
		}
		task.wg.Done()
	}
}

// FillChannelFromTreePool uses a worker pool to walk the tree and fill the channel with its values.
func FillChannelFromTreePool(t *Tree, ch chan int, pool *WorkerPool) {
	var wg sync.WaitGroup
	wg.Add(1)
	pool.tasks <- Task{t, ch, &wg}
	wg.Wait()
	close(ch)
}

func FillSliceFromChannelPool(ch chan int) []int {
	var s []int
	for v := range ch {
		s = append(s, v)
	}
	return s
}
