package main

import (
    "fmt"
    "reflect"
)

func main() {
    var username string = "James"
    peek(username)
}

func peek(v interface{}) {
    fmt.Println("Type: ", reflect.TypeOf(v))
    fmt.Println("Value:", reflect.ValueOf(v))
    fmt.Println("Kind:", reflect.ValueOf(v).Kind())
    fmt.Println("Type: ", reflect.ValueOf(v).Type())
}