package main

import (
    "fmt"
    "reflect"
)

type AliasString string

func main() {
    var username AliasString = "James"
    peek(username)
}

func peek(v interface{}) {
    fmt.Println("Type: ", reflect.TypeOf(v))
    fmt.Println("Value:", reflect.ValueOf(v))
    fmt.Println("Kind:", reflect.ValueOf(v).Kind())
}