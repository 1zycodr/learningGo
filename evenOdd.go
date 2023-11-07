package main

import "fmt"

func main() {
	arr := make([]int, 11, 11)

	for i := 0; i < 11; i++ {
		arr[i] = i
	}

	for _, value := range arr {
		if value%2 == 0 {
			fmt.Printf("%d is even\n", value)
		} else {
			fmt.Printf("%d is odd\n", value)
		}
	}
}
