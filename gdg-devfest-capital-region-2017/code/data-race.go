package main

import (
	"fmt"
	"time"
)

// START data-race OMIT
type counter int

func (c *counter) inc() {
	fmt.Println("pre inc", *c)
	*c++
	fmt.Println("post inc", *c)
	fmt.Println("")
}

func main() {
	c := counter(0)
	go func() { c.inc() }()
	go func() { c.inc() }()
	go func() { c.inc() }()

	time.Sleep(1 * time.Second)
}

// END data-race OMIT
