package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Append doubles slice cap if len+1 was added
	//slice := []int{1, 2, 3}
	//for i := 0; i <= 15; i++ {
	//	slice = append(slice, i)
	//	fmt.Printf("%p, %d, %d\n", slice, cap(slice), len(slice))
	//}
	//
	//slice2 := make([]int, 3, 5)
	//fmt.Println("\nslice2:\n", cap(slice2), len(slice2))
	//for i := 0; i <= 15; i++ {
	//	slice2 = append(slice2, i)
	//	fmt.Printf("%p, %d, %d\n", slice2, cap(slice2), len(slice2))
	//}

	// Append returns pointer on new slice if cap updated
	//slice3 := []int{1, 2, 3, 4}
	//slice3_1 := append(slice3, 4)
	//slice3_1[0] = 11
	//fmt.Println("\nslice3:\n", slice3, "cap:", cap(slice3))
	//fmt.Printf("%p\n", slice3)
	//fmt.Println("\nslice3_1:\n", slice3_1, "cap:", cap(slice3_1))
	//fmt.Printf("%p\n", slice3_1)
	//
	//slice4 := make([]int, 3, 4)
	//slice4_1 := append(slice4, 0)
	//slice4_2 := append(slice4_1, 0)
	//fmt.Println("\nslice4:\n", slice4, "cap:", cap(slice4))
	//fmt.Printf("%p\n", slice4)
	//fmt.Println("\nslice4_1:\n", slice4_1, "cap:", cap(slice4_1))
	//fmt.Printf("%p\n", slice4_1)
	//fmt.Println("\nslice4_2:\n", slice4_2, "cap:", cap(slice4_2))
	//fmt.Printf("%p\n\n", slice4_2)

	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}

	// Slice ranging does not return copy
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceRange1 := mySlice[:3]
	sliceRange2 := mySlice[4:]
	fmt.Println(*(*slice)(unsafe.Pointer(&sliceRange1)))
	fmt.Println(*(*slice)(unsafe.Pointer(&sliceRange2)))
}
