package main

import (
	"fmt"

	gamepkg "github.com/gitshubham45/designPatternGo/snakeAndLadder/gamePkg"
)

func main() {
	fmt.Println("Welcome to Snake and Ladder Game")

	game := gamepkg.Game{
		Positions: make([]*gamepkg.Position, 100),
	}

	for i := 0; i < 100; i++ {
		game.Positions[i] = &gamepkg.Position{
			CurrPos: i + 1,
		}
	}

	fmt.Println("Enter number of Snakes")
	var s int
	fmt.Scanln(&s)

	for i := 1; i <= s; i++ {
		var head int
		var tail int
		fmt.Scan(&head, &tail)
		game.Positions[head-1].InitialPos = head
		game.Positions[head-1].FinalPos = tail
	}

	fmt.Println("Enter number of Ladders")
	var l int
	fmt.Scanln(&l)

	for i := 1; i <= l; i++ {
		var start int
		var end int
		fmt.Scan(&start, &end)
		game.Positions[start-1].InitialPos = start
		game.Positions[start-1].FinalPos = end

		// validate input
		// if start < 1|| start > 100
	}

	fmt.Println("Enter number of Players")
	var p int
	fmt.Scan(&p)

	for i := 1; i <= p; i++ {
		var name string
		fmt.Scan(&name)

		p := &gamepkg.Player{
			Name: name,
			Pos:  1,
		}

		game.Players = append(game.Players, p)

		// validate input
		// if start < 1|| start > 100
	}

}
