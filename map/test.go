package main

import "fmt"

type MapKey struct {
	a int
	b string
}

type MapValue int

func main() {
	m := make(map[MapKey]MapValue)
	k1 := MapKey{1, "1"}
	k2 := MapKey{2, "2"}

	m[k1] = 1
	m[k2] = 2

	fmt.Println(m[k1])
	fmt.Println(m[k2])
}
