package main

import (
	"fmt"
	"strconv"
)

func p1() {
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)
}

func p2() {
	var f float64 = 45.89
	var i int = int(f)
	fmt.Printf("%v\n", i)

}

func p3() {
	var i int = 42
	var s string = strconv.Itoa(i)
	fmt.Printf("%q\n", s)
}

func p4() {
	var s string = "200"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", err, err)
}

func p5() {
	var s string = "200b"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", err, err)
}

func main() {
	p1()
	p2()
	p3()
	p4()
	p5()
}
