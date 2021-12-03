package main

import "fmt"

type Person struct {
	fName string
	lName string
}
type Student struct {
	Person
	stdID int
}

func (p Person) detials() {
	// fmt.Println(p)
	fmt.Println(p.fName + " " + p.lName)
}

func (t Student) details() {
	// fmt.Println(t)
	fmt.Println(t.fName + " " + t.lName)
}

func main() {
	p1 := Person{"Mark", "Sakaberg"}
	p1.detials()

	t1 := Student{Person: Person{"Jeff", "Bezzo"}, stdID: 11}
	t1.details()
}
