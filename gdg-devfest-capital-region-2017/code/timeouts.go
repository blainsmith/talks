package main

import (
	"fmt"
	"strings"
	"time"
)

// START timeouts OMIT
func expensiveTask(password string) chan string {
	result := make(chan string)
	go func() {
		time.Sleep(time.Duration(len(password)) * time.Second)
		result <- strings.ToUpper(password)
		close(result)
	}()
	return result
}

func main() {
	etchan1 := expensiveTask("lu")
	select {
	case hashed := <-etchan1:
		fmt.Println(hashed)
	case <-time.After(3 * time.Second):
		fmt.Println("hashing timed out")

	}
}

// END timeouts OMIT
