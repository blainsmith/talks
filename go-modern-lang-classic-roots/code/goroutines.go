package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintSquare(x int) {
	time.Sleep(2 * time.Second)
	fmt.Println(x * x)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for _, num := range []int{1, 3, 5, 7, 9} {
		go func(num int) {
			PrintSquare(num)
			wg.Done()
		}(num)
	}
	wg.Wait()
}
