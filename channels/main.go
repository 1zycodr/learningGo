package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		value := <-ch
		fmt.Println(value)
	}()

	time.Sleep(1 * time.Second)

	return
}
