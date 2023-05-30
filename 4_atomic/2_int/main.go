package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var count atomic.Int64
	for i := 0; i < 1e5; i++ {
		go func() {
			count.Add(1)
		}()
	}
	time.Sleep(time.Second / 2)
	fmt.Println(count.Load())
}
