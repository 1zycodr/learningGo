package main

import (
	"fmt"
)

func main() {
	testMap := map[int32]int32{
		1: 1,
	}
	for i := 0; i < 100000; i++ {
		testMap[int32(i)] = int32(i)
	}

	for i := 0; i < 100000; i++ {
		a, ok := testMap[int32(i)]
		fmt.Println(a, ok)
	}

	for k, v := range testMap {
		fmt.Println(k, v)
	}
}
