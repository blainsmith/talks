package main

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"time"
)

type idGen struct {
	idchan chan string
}

// START generator-1 OMIT
func NewIDGenerator() *idGen {
	idg := &idGen{
		idchan: make(chan string),
	}
	go func() {
		for {
			b := make([]byte, 1025)
			rand.Read(b)
			idg.idchan <- fmt.Sprintf("%x", sha512.Sum512(b))
		}
	}()
	return idg
}

func (idg *idGen) Next() string {
	return <-idg.idchan
}

func (idg *idGen) Range() <-chan string {
	return idg.idchan
}

// END generator-1 OMIT

// START generator-2 OMIT
func main() {
	ids := NewIDGenerator()

	fmt.Println(ids.Next())
	fmt.Println(ids.Next())
	fmt.Println(ids.Next())

	time.Sleep(2 * time.Second)

	go func() {
		count := 0
		for id := range ids.Range() {
			fmt.Println(count, id)
			count++
		}
	}()

	time.Sleep(1 * time.Second)
}

// END generator-2 OMIT
