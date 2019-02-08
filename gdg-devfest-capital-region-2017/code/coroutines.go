package main

import (
	"fmt"
	"strings"
)

// START coroutines OMIT
func expensiveTask(password string) chan string {
	result := make(chan string)
	go func() {
		result <- strings.ToUpper(password)
		close(result)
	}()
	return result
}

func main() {
	hashed := expensiveTask("luggageCombo12345")
	fmt.Println("Doing other things")
	fmt.Println(<-hashed)
}

// END coroutines OMIT
