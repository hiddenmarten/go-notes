package main

import (
	"sync"
	"time"
)

// WalkNoPool walks the tree t sending all values
// from the tree to the channel ch.
func WalkNoPool(t *Tree, ch chan int, payload time.Duration, wg *sync.WaitGroup) {
	ch <- t.Value
	if t.Left != nil {
		wg.Add(1)
		go WalkNoPool(t.Left, ch, payload, wg)
	}
	if t.Right != nil {
		wg.Add(1)
		go WalkNoPool(t.Right, ch, payload, wg)
	}
	wg.Done()
}

func FillChannelFromTreeNoPool(t *Tree, ch chan int, payload time.Duration) {
	var wg sync.WaitGroup
	wg.Add(1)
	go WalkNoPool(t, ch, payload, &wg)
	wg.Wait()
	close(ch)
}

func FillSliceFromChannelNoPool(ch chan int) []int {
	var s []int
	for v := range ch {
		s = append(s, v)
	}
	return s
}
