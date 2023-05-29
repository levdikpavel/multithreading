package main

import (
	"fmt"
	"time"
)

func main() {
	words := []string{
		"",
		"1",
		"22",
		"333",
		"4444",
	}
	var s string
	go func() {
		for i := 0; i < 1000000; i++ {
			s = words[i%len(words)]
		}
	}()
	go func() {
		for i := 0; i < 1000000; i++ {
			if s != "55555" {
				fmt.Println(s)
			}
		}
	}()
	time.Sleep(time.Second / 2)
}
