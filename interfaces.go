package main

import (
	"fmt"
	"io"
	"os"
)

//type shape interface {
//	getArea() float64
//}
//
//type square struct {
//	a int
//}
//
//func (s square) getArea() float64 {
//	return float64(s.a * s.a)
//}
//
//type triangle struct {
//	base   int
//	height int
//}
//
//func (t triangle) getArea() float64 {
//	return .5 * float64(t.base*t.height)
//}

func main() {
	//resp, err := http.Get("https://www.google.com")
	//if err != nil {
	//	fmt.Println("Error:", err.Error())
	//	os.Exit(0)
	//}
	//defer resp.Body.Close()
	//
	//fmt.Println(resp)
	//
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Error:", err.Error())
	//	os.Exit(1)
	//}
	//
	//fmt.Println(string(body))

	//sq := square{
	//	a: 10,
	//}
	//fmt.Println(sq.getArea())
	//tr := triangle{
	//	height: 10,
	//	base:   10,
	//}
	//fmt.Println(tr.getArea())

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(0)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(0)
	}

	fmt.Println(string(data))
}
