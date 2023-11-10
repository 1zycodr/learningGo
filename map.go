package main

import "fmt"

func main() {
	testMap := map[int]string{
		1: "value 1",
	}

	testMap[2] = "value 2"

	for k, v := range testMap {
		fmt.Println(k, v)
	}

	delete(testMap, 1)

	fmt.Println("After deletion")
	for k, v := range testMap {
		fmt.Println(k, v)
	}
}
