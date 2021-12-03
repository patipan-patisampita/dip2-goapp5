package main

import "fmt"

func main() {
	var value2 interface{}
	value2 = "Welcome"
	show(value2)

	value2 = 49
	show(value2)
}

func show(value2 interface{}) {
	fmt.Printf("%v,%T\n", value2, value2)
}
