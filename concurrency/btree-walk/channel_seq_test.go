package main

import (
	"fmt"
	"testing"
)

var s []int

func BenchmarkInvestigateTreeGeneral(b *testing.B) {

	depths := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	for _, d := range depths {
		t := New(1, d)
		b.Run(fmt.Sprintf("Seq-d%d", d), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ch := make(chan int)
				go FillChannelFromTreeSeq(t, ch)
				s = FillSliceFromChannelSeq(ch)
			}
		})
	}

	for _, d := range depths {
		t := New(1, d)
		b.Run(fmt.Sprintf("NoPool-d%d", d), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ch := make(chan int)
				go FillChannelFromTreeNoPool(t, ch)
				s = FillSliceFromChannelNoPool(ch)
			}
		})
	}

	for _, d := range depths {
		t := New(1, d)
		b.Run(fmt.Sprintf("Pool-d%d", d), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ch := make(chan int)
				pool := NewWorkerPool(poolWorkersCount)
				pool.Start()
				go FillChannelFromTreePool(t, ch, pool)
				s = FillSliceFromChannelPool(ch)
				pool.Stop()
			}
		})
	}
}
