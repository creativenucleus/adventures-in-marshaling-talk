package main

import (
    "fmt"
    "reflect"
)

func main() {
    user := struct {
        FirstName string
    } {
        FirstName: "James",    
    }
    peek(user)
}

func peek(v interface{}) {
    value := reflect.ValueOf(v)
    valueType := value.Type()
    for i := 0; i < value.NumField(); i++ {
        fmt.Printf("Field: %s Type: %s Value: %s\n",
            valueType.Field(i).Name, value.Field(i).Type(), value.Field(i).String())
    }
}