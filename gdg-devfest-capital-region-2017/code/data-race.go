package main

import (
	"fmt"
	"time"
)

// START data-race OMIT
type counter int

func (c *counter) inc() {
	*c++                             // Write data race
	fmt.Println("current value", *c) // Read data race
}

func main() {
	c := counter(0)
	// 3 goroutines reading/writing c
	go func() { c.inc() }()
	go func() { c.inc() }()
	go func() { c.inc() }()

	time.Sleep(1 * time.Second)

	fmt.Println("final value", c)
}

// END data-race OMIT
