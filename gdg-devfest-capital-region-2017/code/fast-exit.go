package main

import (
	"fmt"
	"time"
)

// START fast-exit OMIT
func main() {
	resources := 5

	go func() {
		for i := 1; i <= resources; i++ {
			go func(resource int) {
				fmt.Println("working on resource", resource)
				time.Sleep(2 * time.Second)
				fmt.Println("finished on resource", resource)
			}(i)
		}
	}()
}

// END fast-exit OMIT
