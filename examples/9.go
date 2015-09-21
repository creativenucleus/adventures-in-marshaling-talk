package main

import (
    "fmt"
    "reflect"
)

func main() {
    type User struct {
        TwitterName string `json:"tags go here,yes,here" another:"me,too"`
    }
    
    user := User{
        TwitterName: "@jtruk",
    }
    
    peek(user)
}

func peek(v interface{}) {
    value := reflect.ValueOf(v)
    fmt.Printf("%s", value.Type().Field(0).Tag)
}