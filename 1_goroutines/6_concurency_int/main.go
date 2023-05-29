package main

import (
	"fmt"
	"time"
)

func main() {
	var count int
	for i := 0; i < 1e5; i++ {
		go func() {
			count++
		}()
	}
	time.Sleep(time.Second / 2)
	fmt.Println(count)
}
