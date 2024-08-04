package main

import (
	"time"
)

// WalkSeq walks the tree t sending all values
// from the tree to the channel ch.
func WalkSeq(t *Tree, ch chan int, payload time.Duration) {
	time.Sleep(payload)
	ch <- t.Value
	if t.Left != nil {
		WalkSeq(t.Left, ch, payload)
	}
	if t.Right != nil {
		WalkSeq(t.Right, ch, payload)
	}
}

func FillChannelFromTreeSeq(t *Tree, ch chan int, payload time.Duration) {
	WalkSeq(t, ch, payload)
	close(ch)
}

func FillSliceFromChannelSeq(ch chan int) []int {
	var s []int
	for v := range ch {
		s = append(s, v)
	}
	return s
}
