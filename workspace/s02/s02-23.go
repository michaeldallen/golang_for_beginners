package main

import (
	"fmt"
)

const age = 12
const h_name, h_age = "foo", 20

func f1() {

	const name = "Harry Potter"
	const is_muggle = false
	const age int = 12

	fmt.Printf("%v: %T\n", name, name)
	fmt.Printf("%v: %T\n", is_muggle, is_muggle)
	fmt.Printf("%v: %T\n", age, age)

}

func f2() {
	const pi float64 = 3.14

	var radius float64 = 5.0
	var area float64
	area = pi * radius * radius
	fmt.Printf("Area of a circle with radius %v is %v", radius, area)

}

func main() {
	f1()
	f2()
}
