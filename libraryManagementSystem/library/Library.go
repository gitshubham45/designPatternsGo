package library

import (
	"time"

	"github.com/google/uuid"
)

type Library interface {
	RegisterUser(*User) bool
	// AddBookToLibrary(*Book) bool
	// RemoveBook
	Borrow(string, string) bool
	Return(string, string) bool
}

type LibraryInstance struct {
	ID       string
	Name     string
	Books    map[string]*Book
	Users    map[string]*User
	Requests []*UserRequest
}

func (lib *LibraryInstance) RegisterUser(user *User) bool {
	lib.Users[user.ID] = user
	return true
}

// for noe I am taking ISBN but we can have option for ID as well
func(lib *LibraryInstance) Borrow(bookId string,userId string) bool {
	if _ , ok := lib.Books[bookId] ; ok {
		request := &UserRequest{
			ID : uuid.New().String(),
			UserID: userId,
			BookId: bookId,
			Time: time.Now(),
			RequestType: Borrow,
		}
		lib.Requests = append(lib.Requests, request)
		return true
	}
	return true
}

func(lib *LibraryInstance) Return(ISBN int,userId string) bool {
	return true
}


