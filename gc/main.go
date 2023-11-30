package main

import "runtime"

// Run with: go run -gcflags=-m main.go

func main() {
	var arrayBefore10Mb [1310720]int // меньше 10 мб
	arrayBefore10Mb[0] = 1

	var arrayAfter10Mb [1310721]int // больше 10 мб
	arrayAfter10Mb[0] = 1

	runtime.GC()
}
