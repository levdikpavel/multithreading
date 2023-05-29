package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			for j := 0; j < 10; j++ {
				fmt.Println("i", i, "j", j)
			}
		}(i)
	}

	//wg.Wait()
	time.Sleep(time.Microsecond)
}
