package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 2)
	for t := range ticker.C {
		fmt.Println(t)
	}
}
