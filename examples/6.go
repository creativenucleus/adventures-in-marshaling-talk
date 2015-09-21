package main

import (
    "fmt"
    "reflect"
)

type User struct {
    FirstName string
}

func main() {
    user := User{
        FirstName: "James",
    }
    peek(user)
}

func peek(v interface{}) {
    fmt.Println("Type: ", reflect.TypeOf(v))
    fmt.Println("Value:", reflect.ValueOf(v))
    fmt.Println("Kind:", reflect.ValueOf(v).Kind())
}