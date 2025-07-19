package singleton

import (
	"fmt"
	"sync"
)

// var mu = &sync.Mutex{}
var mu sync.Mutex

type single struct {
}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		// to prevent more than one go routine
		//  bypassing the first check at same time
		if singleInstance == nil {
			fmt.Println("Creating single instance now")
			singleInstance = &single{}
		} else {
			fmt.Println("Single Instance already created")
		}
	} else {
		fmt.Println("Single instance already created")
	}
	return singleInstance
}
