package main

import (
	"fmt"
	"sync"
)

type SplitGroup struct {
	GroupName  string
	Users      []*User
	TotalUsers int
}

func (s *SplitGroup) ShowOne(userID string) {
	fmt.Println("Showing one")
}

func (s *SplitGroup) ShowAll() {
	fmt.Println("Show all")
}

var (
	splitGroupInstance *SplitGroup
	once               sync.Once
)

func NewSplitGroup(n int, name string) *SplitGroup {
	once.Do(func() {
		splitGroupInstance = &SplitGroup{
			GroupName:  name,
			Users:      make([]*User, n),
			TotalUsers: n,
		}
	})

	return splitGroupInstance
}
