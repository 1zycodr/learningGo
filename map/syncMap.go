package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map

	result, ok := sm.Load("hello")
	if ok {
		fmt.Println(result.(string))
	} else {
		fmt.Println("value not found for key: `hello`")
	}

	sm.Store("hello", "world")
	fmt.Println("added value: `world` for key: `hello`")

	result, ok = sm.Load("hello")
	if ok {
		fmt.Printf("result: `%s` found for key: `hello`\n", result.(string))
	}

	ok = sm.CompareAndSwap("hello", "world", "world2")
	if ok {
		fmt.Println("world swapped to world2")
	}

	result, ok = sm.Load("hello")
	if ok {
		fmt.Printf("result: `%s` found for key: `hello`\n", result.(string))
	}
}
