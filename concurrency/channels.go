package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("started 1")
		ch <- 1 // blocking
		fmt.Println("writed 1")
	}()

	go func() {
		fmt.Println("started 2")
		ch <- 2 // blocking
		fmt.Println("writed 2")
	}()

	time.Sleep(1 * time.Second)

	fmt.Printf("received first: %d\n", <-ch)
	fmt.Printf("received second: %d\n", <-ch)
}
