package main

import (
	"fmt"

	"github.com/gitshubham45/designPatternGo/singleton/1/singleton"
)

func main() {

	// for i := 0; i < 30; i++ {
	// 	go getInstance()
	// }

	for range 30 {
		go singleton.GetInstance()
		// fmt.Println(s)
	}

	fmt.Scanln()
}
