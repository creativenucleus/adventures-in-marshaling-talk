package main

import (
    "encoding/json"
    "os"
)

type User struct {
    FirstName string
    PasswordHash string
}

func main() {    
    user := User{
        FirstName: "James",
        PasswordHash: "P455w0rd!",
    }

    out, _ := json.MarshalIndent(user, "", "   ")
        
    os.Stdout.Write(out)
}