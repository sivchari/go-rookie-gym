package group

type Group struct {
	ID     int
	UserID int
	Name   string
}

func NewGroup(userID int, name string) *Group {
	return &Group{
		UserID: userID,
		Name:   name,
	}
}
