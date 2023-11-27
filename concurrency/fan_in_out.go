package main

import (
	"fmt"
	"sync"
)

func numbersGenerator(num int) <-chan int {
	ch := make(chan int, num)

	go func() {
		for i := 1; i <= num; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func isNumberEven(numbers <-chan int) <-chan string {
	res := make(chan string)
	go func() {
		for value := range numbers {
			result := "number %d is odd"
			if value%2 == 0 {
				result = "number %d is even"
			}
			res <- fmt.Sprintf(result, value)
		}
		close(res)
	}()
	return res
}

func merge(inputs ...<-chan string) <-chan string {
	res := make(chan string)
	wg := sync.WaitGroup{}

	for _, inp := range inputs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for value := range ch {
				res <- value
			}
		}(inp)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func main() {
	gen := numbersGenerator(100)

	// fan-out - разделяем один канал на несколько
	ch1 := isNumberEven(gen)
	ch2 := isNumberEven(gen)
	ch3 := isNumberEven(gen)

	// fan-in - соединяем все каналы в один
	res := merge(ch1, ch2, ch3)

	// выводим результат с канала
	for result := range res {
		fmt.Println(result)
	}
}
