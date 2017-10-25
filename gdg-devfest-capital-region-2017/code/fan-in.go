package main

import (
	"fmt"
	"sync"
	"time"
)

// START fan-in-1 OMIT
func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// END fan-in-1 OMIT

// START fan-in-2 OMIT
func combine(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(len(cs))
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// END fan-in-2 OMIT

// START fan-in-3 OMIT
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	in := gen(nums...)

	sq1 := sq(in)
	sq2 := sq(in)
	sq3 := sq(in)

	for result := range combine(sq1, sq2, sq3) {
		fmt.Println(result)
	}

	time.Sleep(1 * time.Second)
}

// END fan-in-3 OMIT
