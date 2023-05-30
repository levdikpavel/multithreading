package main

import "fmt"

func main() {
	var a chan int
	//a <- 1
	//fmt.Println(<-a)
	select {
	case i := <-a:
		fmt.Println(i)
	default:
		fmt.Println("default")
	}
	close(a)
}
