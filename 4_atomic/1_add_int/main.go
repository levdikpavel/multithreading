package main

import (
	"fmt"
	"time"
)

func main() {
	var count int64
	for i := 0; i < 1e5; i++ {
		go func() {
			count++
			//atomic.AddInt64(&count, 1)
		}()
	}
	time.Sleep(time.Second / 2)
	fmt.Println(count)
}
