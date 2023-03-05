package main

import "fmt"

func main() {

	var name string
	name = "fubar"
	i := 55

	i = 66

	fmt.Printf("Template string %q, next %d\n\n", name, i)

	var (
		s  string = "foo"
		i2 int    = 5
	)

	fmt.Printf("foo: %s, bar: %d", s, i2)

}
