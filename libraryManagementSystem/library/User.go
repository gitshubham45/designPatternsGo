package library

type User struct {
	ID       string
	Name     string
	Phone    string
	Requests []*UserRequest
}
