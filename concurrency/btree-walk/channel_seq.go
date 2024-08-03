package main

import (
	"time"
)

// WalkSeq walks the tree t sending all values
// from the tree to the channel ch.
func WalkSeq(t *Tree, ch chan int) {
	time.Sleep(pseudoPayload)
	ch <- t.Value
	if t.Left != nil {
		WalkSeq(t.Left, ch)
	}
	if t.Right != nil {
		WalkSeq(t.Right, ch)
	}
}

func FillChannelFromTreeSeq(t *Tree, ch chan int) {
	WalkSeq(t, ch)
	close(ch)
}

func FillSliceFromChannelSeq(ch chan int) []int {
	var s []int
	for v := range ch {
		s = append(s, v)
	}
	return s
}
