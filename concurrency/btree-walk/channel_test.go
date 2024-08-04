package main

import (
	"fmt"
	"testing"
	"time"
)

var s []int

func BenchmarkInvestigateTreeGeneral(b *testing.B) {

	depths := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}
	payloads := []time.Duration{2 * time.Millisecond, 4 * time.Millisecond,
		8 * time.Millisecond, 16 * time.Millisecond, 32 * time.Millisecond,
		64 * time.Millisecond, 128 * time.Millisecond, 256 * time.Millisecond,
		512 * time.Millisecond, 1024 * time.Millisecond, 2048 * time.Millisecond,
	}

	// Way too long
	//for _, d := range depths {
	//	for _, p := range payloads {
	//		t := New(1, d)
	//		b.Run(fmt.Sprintf("Seq-d%d-p%d", d, p/1000000), func(b *testing.B) {
	//			for i := 0; i < b.N; i++ {
	//				ch := make(chan int)
	//				go FillChannelFromTreeSeq(t, ch, p)
	//				s = FillSliceFromChannelSeq(ch)
	//			}
	//		})
	//	}
	//}

	for _, d := range depths {
		for _, p := range payloads {
			t := New(1, d)
			b.Run(fmt.Sprintf("NoPool-d%d-p%d", d, p/1000000), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ch := make(chan int)
					go FillChannelFromTreeNoPool(t, ch, p)
					s = FillSliceFromChannelNoPool(ch)
				}
			})
		}
	}

	for _, d := range depths {
		for _, p := range payloads {
			t := New(1, d)
			b.Run(fmt.Sprintf("Pool-d%d-p%d", d, p/1000000), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ch := make(chan int)
					pool := NewWorkerPool(poolWorkersCount)
					pool.Start()
					go FillChannelFromTreePool(t, ch, p, pool)
					s = FillSliceFromChannelPool(ch)
					pool.Stop()
				}
			})
		}
	}
}
