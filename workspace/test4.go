package main

import "fmt"

func main() {
	var b bool
	var i int
	var f float64
	var s string

	fmt.Printf("%s: %T, %t\n", "b", b, b)
	fmt.Printf("%s: %T, %d\n", "i", i, i)
	fmt.Printf("%s: %T, %f\n", "f", f, f)
	fmt.Printf("%s: %T, %q\n", "s", s, s)
}
