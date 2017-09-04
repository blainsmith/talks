package main

import (
	"log"
	"time"
)

func πer(in chan string) {
	for word := range in {
		log.Println(word + " PIE!!!")
	}
}

func main() {
	in := make(chan string)

	go πer(in)

	for _, word := range []string{"apple", "blueberry", "pumpkin"} {
		in <- word
	}

	time.Sleep(2 * time.Second)
}
