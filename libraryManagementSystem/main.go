package main

import (
	"fmt"

	"github.com/gitshubham45/designPatternGo/libraryManagementSystem/library"
)

func main() {

	fmt.Println("Welcome to Central Library")

	lib := library.NewLibraryInstance("Central Library")
	library.LoadBooksFromStore("./library/books.json", lib)
	fmt.Println("Enter your phone number to login/Register")

	for {
		var phone string

		fmt.Scanln(&phone)

		_, ok := lib.FindUser(phone)
		if !ok {
			_ = lib.RegisterUser(phone)
			fmt.Println("User Registered")
		} else {
			fmt.Println("User Logged in")
		}

		var bookID string
		var requestType library.RequestType
		fmt.Println("Enter book ID")
		fmt.Scanln(&bookID)

		var req string
		fmt.Println("Enter request Type: [Borrow , Return]")
		fmt.Scanln(&req)

		switch req {
		case "Borrow":
			requestType = library.Borrow
		case "Return":
			requestType = library.Return
		default:
			fmt.Println("Invalid request type")
			continue
		}

		lib.FullFillRequest(bookID, phone, requestType)

		if requestType == library.Borrow {
			fmt.Printf("Book borrowed, BookID : %s", bookID)
			continue
		} else if requestType == library.Return {
			fmt.Printf("Book returned, BookID : %s", bookID)
			continue
		}
	}
}
