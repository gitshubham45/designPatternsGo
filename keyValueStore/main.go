package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gitshubham45/designPatternGo/keyValueStore/store"
)

func main() {

	myStore := store.NewKeyValueStore()

	args := os.Args

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter commands (type 'exit' to quit)")

	for {
		if len(args) < 1 {
			continue
		}
		
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		parts := strings.Fields(input) // split by space
		if len(parts) == 0 {
			continue
		}

		command := strings.ToUpper(parts[0])
		args := parts[1:]

		switch command {
		case "SET":
			if len(args) < 2 {
				fmt.Println("Usage : SET key val <time[s,h,m,d](optional)> ")
			}
			if len(args) == 2 {
				myStore.SET(args[0], args[1], nil)
			} else {
				myStore.SET(args[0], args[1], &args[2])
			}
		case "GET":
			output := myStore.GET(args[0])
			fmt.Printf("Key : %s - Val : %s", args[0], output)
		case "DEL":
			ok := myStore.DEL(args[0])
			if !ok {
				fmt.Println("Key is not found")
			}

		}

	}
}
