package main

import "fmt"

func main() {
	var city string = "London"
	country := "USA"
	{
		country := "UK"
		fmt.Println(country)
		fmt.Println(city)
	}
	fmt.Println(country)
	fmt.Println(city)
}
