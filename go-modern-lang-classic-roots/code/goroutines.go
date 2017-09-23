package main

import (
	"fmt"
	"time"
)

func PrintSquare(x int) {
	time.Sleep(2 * time.Second)
	fmt.Println(x * x)
}

func main() {
	for _, num := range []int{1, 3, 5, 7, 9} {
		go PrintSquare(num)
	}

	// Force the main() goroutine to wait until all others finish
	time.Sleep(5 * time.Second)
}
