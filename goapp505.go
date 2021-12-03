package main

import "fmt"

type vehicle interface {
	accelerate()
}

func testVehicle(v vehicle) {
	fmt.Println(v)
}

type Car struct {
	model string
	color string
}

func (c Car) accelerate() {
	fmt.Println("Acelrating")
}

type Toyota struct{
	model string
	color string
	speed int
}
func (t Toyota)accelerate(){
	fmt.Println("I am toyota, I accelerate fast?")
}

func main() {
	c1 := Car{"Toyota", "Blue"}
	c1.accelerate()
	testVehicle(c1)
	
	t1:= Toyota{"Toyota", "Red", 100}
	t1.accelerate()
	testVehicle(t1)
}
