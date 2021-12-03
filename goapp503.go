package main

import (
	"fmt"
	"strconv"
)

//Defining type struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//Method (value reciver)
func (p Person) hello() string {
	return p.firstName + " " + p.lastName + " " + p.city + " " + p.gender + " " + strconv.Itoa((p.age))
}
func main() {
	//initialization
	p1 := Person{firstName: "Mark", lastName: "Sakaberg", city: "Trat", gender: "Male", age: 25}
	p2 := Person{"Jeff", "Bezzo", "Montona", "Male", 35}
	p3 := Person{firstName: "Steav", lastName: "Job"}
	fmt.Println(p1.hello())
	fmt.Println(p2.hello())
	fmt.Println(p3.hello())
}
