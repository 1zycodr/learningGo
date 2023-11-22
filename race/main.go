package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "1"

	var mutex sync.Mutex

	wg.Add(4)
	go updateMessage("2", &mutex)
	go updateMessage("3", &mutex)
	go updateMessage("4", &mutex)
	go updateMessage("5", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
