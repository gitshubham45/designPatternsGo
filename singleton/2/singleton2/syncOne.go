package singleton2

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("creating single instance now")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("single instance already created.")
	}

	return singleInstance
}
