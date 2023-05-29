package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("hello from anonymous func")
	}()

	go printHello()

	var p printer
	go p.printHello()

	time.Sleep(1)
}

func printHello() {
	fmt.Println("hello from named func")
}

type printer struct {
}

func (printer) printHello() {
	fmt.Println("hello from struct method")
}
