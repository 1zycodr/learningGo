package main

import "fmt"

func main() {
	// The size of an array is part of its type.
	// The types [10]int and [20]int are distinct.
	arr := [3]int{1, 2, 3}

	// Arrays are values. Assigning one array to another copies all the elements.
	arr2 := arr
	fmt.Printf("%p, %p\n", &arr, &arr2)

	// If you pass an array to a function,
	// it will receive a copy of the array, not a pointer to it
	arrReceiver(arr)
}

func arrReceiver(arr [3]int) {
	fmt.Printf("%p", &arr)
}
