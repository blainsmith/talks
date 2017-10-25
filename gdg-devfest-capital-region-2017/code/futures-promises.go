package main

import (
	"fmt"
	"time"
)

// START main OMIT
func expensiveComputeSquare(num int) chan int {
	computed := make(chan int)
	go func(num int) {
		time.Sleep(time.Duration(num) * time.Second)
		computed <- num * num
	}(num)
	return computed
}

func addSquares(a int, b int, c int) int {
	ecs1 := expensiveComputeSquare(a)
	esc2 := expensiveComputeSquare(b)
	esc3 := expensiveComputeSquare(c)
	return <-ecs1 + <-esc2 + <-esc3 // Blocks on returning until all receive a value
}

func main() {
	sum1Start := time.Now()
	sum1 := addSquares(2, 3, 1)
	fmt.Printf("%d (took %d seconds)\n", sum1, time.Since(sum1Start)/time.Second)
	sum2Start := time.Now()
	sum2 := addSquares(3, 5, 1)
	fmt.Printf("%d (took %d seconds)\n", sum2, time.Since(sum2Start)/time.Second)
}

// END main OMIT
