package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ch1 := make(chan string, 1)
	ch2 := make(chan int, 1)

	go DoStuff1(ch1)
	go DoStuff2(ch2)

	res1 := <-ch1
	res2 := <-ch2

	fmt.Printf("stuff1: %s\nstuff2: %d\n", res1, res2)
	fmt.Println(time.Since(start))
}

func DoStuff1(ch chan string) {
	time.Sleep(100 * time.Millisecond)
	ch <- "do stuff 1 finished"
}

func DoStuff2(ch chan int) {
	time.Sleep(150 * time.Millisecond)
	ch <- 2
}
