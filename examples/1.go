// NB Not a real program!

import (
    "encoding/json"
    "encoding/xml"
)

type User struct {
    FirstName string `json:"first_name,omitempty" xml:"omitempty"`
    PasswordHash string `json:"-" xml:"-"`
}

user := User{
    FirstName: "James",
    PasswordHash: "P455w0rd!",
}

json, err := json.Marshal(user)

xml, err := xml.Marshal(user)
