package main
import (
    "fmt"
    "reflect"
)



func p1() {
    var grades int = 42
    var message string = "hello world"
    var isCheck bool = true
    var amount float32 = 5466.54
    
    fmt.Printf("variable grades == %v is of type %T\n", grades, grades)
    fmt.Printf("variable message == %v is of type %T\n", message, message)
    fmt.Printf("variable isCheck == %v is of type %T\n", isCheck, isCheck)
    fmt.Printf("variable amount == %v is of type %T\n", amount, amount)
    
}


func p2() {
    fmt.Printf("Type: %v\n", reflect.TypeOf(1000))
    fmt.Printf("Type: %v\n", reflect.TypeOf("foo"))
    fmt.Printf("Type: %v\n", reflect.TypeOf(46.0))
    fmt.Printf("Type: %v\n", reflect.TypeOf(46))
    fmt.Printf("Type: %v\n", reflect.TypeOf(true))
    
}


func p3() {
    var grades int = 99
    var message string = "Henry and Finnian are awesome"
    
    fmt.Printf("variable grades == %v is of type %v\n", grades, reflect.TypeOf(grades))
    fmt.Printf("variable message == %v is of type %v\n", message, reflect.TypeOf(message))

    fmt.Printf("type of refleted type is %v\n", reflect.TypeOf(reflect.TypeOf("foo")))
}

func main() {
    p1()
    p2()
    p3()
}