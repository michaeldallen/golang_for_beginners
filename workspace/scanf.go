package main

import "fmt"

func main() {

	var name string
	var is_muggle bool

	fmt.Print("Enter your name & muggleness: ")
	fmt.Scanf("%s %t", &name, &is_muggle)
	fmt.Printf("hi %q (%t)\n", name, is_muggle)

}
