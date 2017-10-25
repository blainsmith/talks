package main

import (
	"fmt"
	"time"
)

// START fan-out-1 OMIT
func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func sq(in <-chan int) {
	go func() {
		for n := range in {
			fmt.Println(n * n)
		}
	}()
}

// END fan-out-1 OMIT

// START fan-out-2 OMIT
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	in := gen(nums...)

	sq(in)
	sq(in)
	sq(in)

	time.Sleep(1 * time.Second)
}

// END fan-out-2 OMIT
