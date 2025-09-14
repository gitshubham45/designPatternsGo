package pkg

type User struct {
	UserName    string
	Name        string
	Email       string
	PhoneNumber string
	Index       int
}

func NeUser(userName, name, email, phoneNumber string, index int) *User {
	return &User{
		UserName:    userName,
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		Index:       index,
	}
}
