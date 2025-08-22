package library

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Library interface {
	FindUser(string) string
	RegisterUser(*User) bool
	AddBookToLibrary(*Book) bool
	// ListBoos
	// RemoveBook
	FullFillRequest(string, string, RequestType) bool
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

func (lib *LibraryInstance) FindUser(phone string) (string, bool) {
	if _, present := lib.Users[phone]; !present {
		fmt.Println("User is not registered, Please register")
		return "false", false
	}

	userID := lib.Users[phone].ID

	return userID, true
}

func (lib *LibraryInstance) RegisterUser(phone string)  bool {
	newuser := &User{
		ID:       uuid.New().String(),
		Phone:    phone,
		Requests: make([]*UserRequest, 0),
	}
	lib.Users[phone] = newuser
	return  true
}

func (lib *LibraryInstance) AddBookToLibrary(book *Book) bool {
	if _, present := lib.Books[book.ID]; !present {
		lib.Books[book.ID] = book
		return true
	}
	return false
}

// for noe I am taking ISBN but we can have option for ID as well
func (lib *LibraryInstance) FullFillRequest(bookID string, userID string, requestType RequestType) bool {
	if _, ok := lib.Books[bookID]; ok {
		request := &UserRequest{
			ID:          uuid.New().String(),
			UserID:      userID,
			BookId:      bookID,
			Time:        time.Now(),
			RequestType: requestType,
		}
		lib.Requests = append(lib.Requests, request)
		lib.Users[userID].Requests = append(lib.Users[userID].Requests, request)
		return true
	}
	return true
}

var (
	libraryInstance *LibraryInstance
	once            sync.Once
)

func NewLibraryInstance(name string) *LibraryInstance {
	once.Do(func() {
		libraryInstance = &LibraryInstance{
			ID:       uuid.New().String(),
			Name:     name,
			Books:    make(map[string]*Book),
			Users:    make(map[string]*User),
			Requests: make([]*UserRequest, 7),
		}
	})
	return libraryInstance
}
