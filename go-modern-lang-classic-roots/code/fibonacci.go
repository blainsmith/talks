package main

import (
	"fmt"
)

func fibonacci(n int, c chan int64) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- int64(x)
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int64, 90)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
}
