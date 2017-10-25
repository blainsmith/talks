package main

import (
	"fmt"
	"sync"
	"time"
)

// START waitgroup OMIT
func main() {
	resources := 5
	var wg sync.WaitGroup
	wg.Add(resources)

	go func() {
		for i := 1; i <= resources; i++ {
			go func(resource int) {
				fmt.Println("working on resource", resource)
				time.Sleep(2 * time.Second)
				fmt.Println("finished on resource", resource)
				wg.Done()
			}(i)
		}
	}()

	wg.Wait()
}

// END waitgroup OMIT
