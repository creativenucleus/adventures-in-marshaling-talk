type User struct {
    FirstName string `json:"first_name,omitempty" xml:"omitempty"`
    PasswordHash string `json:"-" xml:"-"`
}
