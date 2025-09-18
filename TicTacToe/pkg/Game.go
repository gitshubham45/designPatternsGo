package pkg

import (
	"fmt"
	"strconv"
)

type Game struct {
	Grid [3][3]string
}

func checkWinner(grid *[3][3]string, mark string) bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if grid[i][0] == mark && grid[i][1] == mark && grid[i][2] == mark {
			return true
		}
	}

	// Check columns
	for j := 0; j < 3; j++ {
		if grid[0][j] == mark && grid[1][j] == mark && grid[2][j] == mark {
			return true
		}
	}

	// Check diagonals
	if grid[0][0] == mark && grid[1][1] == mark && grid[2][2] == mark {
		return true
	}
	if grid[0][2] == mark && grid[1][1] == mark && grid[2][0] == mark {
		return true
	}

	// No winner
	return false
}


// func checkWinner(grid *[3][3]string, mark string, x, y int) bool {
// 	for i := 0; i < 3; i++ {
// 		if grid[i][0] == mark && grid[i][0] == grid[i][1] && grid[i][1] == grid[i][2] {
// 			return true
// 		}
// 	}

// 	// Check columns
// 	for j := 0; j < 3; j++ {
// 		if grid[0][j] == mark && grid[0][j] == grid[1][j] && grid[1][j] == grid[2][j] {
// 			return true
// 		}
// 	}

// 	// Check diagonals
// 	if grid[0][0] == mark && grid[0][0] == grid[1][1] && grid[1][1] == grid[2][2] {
// 		return true
// 	}
// 	if grid[0][2] != "" && grid[0][2] == grid[1][1] && grid[1][1] == grid[2][0] {
// 		return true
// 	}

// 	// No winner
// 	return false
// }

func (g *Game) MakeMove(parts []string, p *Player) (bool, error) {
	x, err := strconv.Atoi(parts[0])

	if err != nil {
		fmt.Println("Error converting to string- check the input again")
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Error converting to string- check the input again")
	}

	g.Grid[x][y] = p.Mark

	return checkWinner(&g.Grid, p.Mark), nil
}
