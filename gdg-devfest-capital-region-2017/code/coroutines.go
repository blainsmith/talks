package main

import (
	"fmt"
	"strings"
)

// START coroutines OMIT
func expensiveTask(password string) chan string {
	result := make(chan string)
	err := make(chan error)
	go func() {
		result <- strings.ToUpper(password)
		close(result)
	}()
	return result, err
}

func main() {
	hashed, err := expensiveTask("luggageCombo12345")
	fmt.Println("Doing other things")
	fmt.Println(<-hashed)
}

// END coroutines OMIT
