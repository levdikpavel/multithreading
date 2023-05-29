package main

import "fmt"

func main() {
	c := make(chan int)
	generate(c)

	for x := range c {
		fmt.Println(x)
	}
}

func generate(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}
