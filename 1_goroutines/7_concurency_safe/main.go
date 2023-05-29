package main

import (
	"fmt"
	"time"
)

func main() {
	arr := make([]int, 1e6)
	fmt.Println(len(arr), arr[:5], arr[len(arr)-5:])

	for i := 0; i < len(arr); i++ {
		go func(i int) {
			arr[i] = i * i
		}(i)
	}

	time.Sleep(time.Second / 2)
	fmt.Println(len(arr), arr[:5], arr[len(arr)-5:])
}
