package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- rand.Intn(10)
		}
		close(ch)
	}()

	for value := range ch {
		fmt.Printf("received value: %d\n", value)
	}
}
