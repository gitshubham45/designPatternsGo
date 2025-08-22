package library

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadBooksFromStore(filename string, lib *LibraryInstance) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", err)
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file", err)
		return err
	}

	var books []*Book
	if err := json.Unmarshal(bytes, &books); err != nil {
		fmt.Println("Error unmarshalling file" , err)
		return err
	}

	for _, book := range books {
		fmt.Println(book)
		lib.AddBookToLibrary(book)
	}

	fmt.Println("Books loaded to library")

	return nil
}
