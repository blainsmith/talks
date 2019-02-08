package main

import (
	"fmt"
	"sync"
	"time"
)

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

func dbl(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

func mult(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 10000
		}
		close(out)
	}()
	return out
}

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

// START pipeline OMIT
func main() {
	evens := gen(2, 4, 6, 8, 10)
	odds := gen(1, 3, 5, 7, 9)

	sqe1 := sq(evens) // 3 routines to square even numbers
	sqe2 := sq(evens)
	sqe3 := sq(evens)
	sqo1 := sq(odds) // 2 routines to square odd numbers
	sqo2 := sq(odds)
	csq := combine(sqe1, sqe2, sqe3, sqo1, sqo2) // collect the square results
	dbl1 := dbl(csq)                             // 2 routines to double the results
	dbl2 := dbl(csq)
	cdbl := combine(dbl1, dbl2) // collect the double resultscd
	mul1 := mult(cdbl)
	mul2 := mult(cdbl)
	mul3 := mult(cdbl)

	for result := range combine(mul1, mul2, mul3) {
		fmt.Println(result)
	}

	time.Sleep(1 * time.Second)
}

// END pipeline OMIT
