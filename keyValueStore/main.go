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

	// for {
	// 	var command string
	// 	var key string
	// }
	args := os.Args
	fmt.Println(args[1])

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter commands (type 'exit' to quit)")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		parts := strings.Fields(input) // split by space
		if len(parts) == 0 {
			continue
		}

		command := parts[0]
		args := parts[1:]

		switch command {
		case "SET":
			if len(args) < 2 {
				fmt.Println("Usage : SET key val <time(optional)> ")
			}
			if len(args) == 2 {
				myStore.SET(args[0],args[1])
			}else{
				myStore.SET(args[0],args[1],args[2])
			}
			
		}

	}
}
