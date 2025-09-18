package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gitshubham45/designPatternGo/TicTacToe/pkg"
)

func printGrid(grid *[3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(grid[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Welcome to Tic Tac Toe game")

	game := pkg.Game{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game.Grid[i][j] = "-"
		}
	}

	fmt.Print(">")
	fmt.Println("Enter the name of Player 1 - (0) :")
	var p1 string
	fmt.Scanln(&p1)
	fmt.Println("Enter the name of Player 1 - (X) :")
	var p2 string
	fmt.Scanln(&p2)

	u1 := &pkg.Player{
		Name: p1,
		Mark: "0",
	}

	u2 := &pkg.Player{
		Name: p2,
		Mark: "X",
	}

	fmt.Println(u1, u2)

	printGrid(&game.Grid)

	fmt.Println("Game Started...")

	moveP1 := true

	scanner := bufio.NewScanner(os.Stdin)

	for {
		var currentPlayer *pkg.Player
		if moveP1 {
			currentPlayer = u1
		} else {
			currentPlayer = u2
		}
		fmt.Printf("%s(%s) Enter Your next move by selectig grid Postion", currentPlayer.Name, currentPlayer.Mark)

		fmt.Print(">")
		scanner.Scan()

		input := scanner.Text()

		parts := strings.Fields(input)

		fmt.Println("Parts" , parts)

		if len(parts) > 2 {
			fmt.Println("Input of invalid length")
		}

		playerWon, err := game.MakeMove(parts, currentPlayer)
		if err != nil {
			fmt.Println("Error ")
		}

		if playerWon {
			fmt.Printf("%s Won\n", currentPlayer.Name)
			fmt.Println("Start New game")
			break
		} else {
			moveP1 = !moveP1	
		}
		printGrid(&game.Grid)
	}
}
