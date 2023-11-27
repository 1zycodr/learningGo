package main

import (
	"fmt"
	"time"
)

func longTask(duration int, result chan string) {
	time.Sleep(time.Duration(duration) * time.Second)
	result <- fmt.Sprintf("result after %ds. duration", duration)
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go longTask(5, ch1)
	go longTask(5, ch2)
	go longTask(5, ch3)

	select {
	case res := <-ch1:
		fmt.Println("Result 1:", res)
	case res := <-ch2:
		fmt.Println("Result 2:", res)
	case res := <-ch3:
		fmt.Println("Result 3:", res)
	case <-time.After(6 * time.Second):
		fmt.Println("No response received from tasks")
	}
}
