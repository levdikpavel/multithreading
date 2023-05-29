package main

import (
	"time"
)

func main() {
	m := make(map[int]int)

	go func() {
		for i := 0; i < 100000; i++ {
			m[0] = i
		}
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			if m[0] < 0 {
				// some action
			}
		}
	}()

	time.Sleep(time.Second / 2)
}
