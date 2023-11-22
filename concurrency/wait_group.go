package main

import (
	"fmt"
	"sync"
)

func printString(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

//func main() {
//	var wg sync.WaitGroup
//
//	words := []string{
//		"a", "b", "c", "d", "e", "f",
//	}
//
//	for _, word := range words {
//		go printString(word, &wg)
//		wg.Add(1)
//	}
//
//	wg.Wait()
//}
