package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	var s atomic.Value
	s.Store(&stats{
		m:   map[int]int{0: 42},
		arr: []int{42},
	})

	go func() {
		for i := 0; i < 100000; i++ {
			s.Store(&stats{
				m:   map[int]int{0: i},
				arr: []int{i},
			})
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			v := s.Load().(*stats)
			fmt.Println(v.m[0], v.arr[0])
			runtime.Gosched()
		}
	}()

	time.Sleep(time.Second / 2)
}

type stats struct {
	m   map[int]int
	arr []int
}
