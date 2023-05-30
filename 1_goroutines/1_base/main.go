package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("hello from anonymous func")
		//time.Sleep(time.Second * 2)
		//fmt.Println("hello 2")
	}()

	go printHello()

	var p printer
	go p.printHello()

	time.Sleep(time.Second)
	//fmt.Scanln()
}

func printHello() {
	fmt.Println("hello from named func")
}

type printer struct {
}

func (printer) printHello() {
	fmt.Println("hello from struct method")
}
