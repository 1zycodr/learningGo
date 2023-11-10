package main

import "fmt"

type contact struct {
	email string
	phone string
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	amirlan := person{
		firstName: "Amirlan",
		lastName:  "Mukhitdinov",
		contact: contact{
			email: "test",
			phone: "+77072402557",
		},
	}
	fmt.Printf("%+v", amirlan)
}
