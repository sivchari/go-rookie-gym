package user

type User struct {
	ID   string
	Name string
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}
