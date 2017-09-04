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
		PrintSquare(num)
	}
}
