package main

import (
	"fmt"
	"sync"
	"time"
)

// START mutex OMIT
var mu sync.RWMutex

type counter int

func (c *counter) inc() {
	mu.Lock()                        // Lock the mutex from others
	*c++                             // Increment c
	fmt.Println("current value", *c) // Print c
	mu.Unlock()                      // Unlock mutex so others can read/write
}

func main() {
	c := counter(0)
	// 3 goroutines reading/writing c
	go func() { c.inc() }()
	go func() { c.inc() }()
	go func() { c.inc() }()

	time.Sleep(1 * time.Second)

	mu.RLock()
	fmt.Println("final value", c)
	mu.RUnlock()
}

// END mutex OMIT
