package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan struct{})
	go func() {
		fastTicker := time.NewTicker(time.Second * 5 / 2)
		for _ = range fastTicker.C {
			select {
			case c <- struct{}{}:
			default:
				fmt.Println("skip fast job")
			}
		}
	}()

	time.Sleep(time.Second / 2)
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			fmt.Println("---------- long job", time.Since(start))
			time.Sleep(time.Second * 4)
		case <-c:
			fmt.Println("---------- fast job", time.Since(start))
		}
	}
}
