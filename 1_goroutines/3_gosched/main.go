package main

import (
	"fmt"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(1)
	for i := 0; i < 3; i++ {
		i := i
		go func() {
			for j := 0; j < 10; j++ {
				fmt.Println("i", i, "j", j)
				//runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Microsecond)
}
