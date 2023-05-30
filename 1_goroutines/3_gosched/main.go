package main

import (
	"fmt"
	"time"
)

func main() {
	demo1()
	//demo2()
}

func demo1() {
	//runtime.GOMAXPROCS(1)
	for i := 0; i < 3; i++ {
		i := i
		go func() {
			// 10000
			for j := 0; j < 10; j++ {
				fmt.Println("i", i, "j", j)
				//runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Microsecond)
	fmt.Scanln()
}

func demo2() {
	var n int
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				current := n + 1
				//runtime.Gosched()
				n = current
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(n)
}
