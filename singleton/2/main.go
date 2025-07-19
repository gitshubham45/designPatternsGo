package main

import (
	"fmt"

	"github.com/gitshubham45/designPatternGo/singleton/2/singleton2"
)

func main() {

	for range 30 {
		// s := singleton2.GetInstance()
		// fmt.Println(s)

		
		go singleton2.GetInstance() // output - creating single instance now (only)
		// all go routines are going inside the nil check condtion at same time
	}

	// time.Sleep(5*time.Second)

	fmt.Scanln()
}
