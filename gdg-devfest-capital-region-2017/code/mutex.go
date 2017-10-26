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
	mu.Lock() // Lock the mutex from others
	fmt.Println("pre inc", *c)
	*c++
	fmt.Println("post inc", *c)
	fmt.Println("")
	mu.Unlock() // Unlock mutex so others can read/write
}

func main() {
	c := counter(0)
	go func() { c.inc() }()
	go func() { c.inc() }()
	go func() { c.inc() }()

	time.Sleep(1 * time.Second)
}

// END mutex OMIT
