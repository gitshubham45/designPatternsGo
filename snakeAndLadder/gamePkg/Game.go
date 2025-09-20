package gamePkg

type Position struct {
	CurrPos    int
	InitialPos int
	FinalPos   int
}

type Game struct {
	Positions []*Position
	Players   []*Player
}

// snake
