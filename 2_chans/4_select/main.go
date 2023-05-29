package main

import "fmt"

func main() {
	c1 := make(chan int, 5)
	c2 := make(chan int, 5)
	for i := 1; i <= 3; i++ {
		c1 <- i
		c2 <- -i
	}
	close(c1)
	close(c2)

	for i := 1; i <= 10; i++ {
		select {
		case x := <-c1:
			fmt.Println("c1:", x)
		case x := <-c2:
			fmt.Println("c2:", x)
		default:
			fmt.Println("NO DATA!")
		}
	}
}
