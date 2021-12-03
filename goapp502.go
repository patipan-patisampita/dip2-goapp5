package main

import "fmt"

//Defining type struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//Accessing Struct
func main() {
	var p4 Person
	p1 := Person{firstName: "Mark", lastName: "Sakaberg", city: "Trat", gender: "Male", age: 25}
	p2 := Person{"Jeff", "Bezzo", "Montona", "Male", 35}
	p3 := Person{firstName: "Steav", lastName: "Job"}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p1.firstName, " ", p1.lastName)

	p4.firstName = "Luca"
	p4.lastName = "Lili"
	p4.city = "Toronto"
	p4.gender = "Female"
	p4.age = 25

	printPerson_4(p4)

}
func printPerson_4(per_4 Person ){

	fmt.Println(per_4.firstName, per_4.lastName)
}
