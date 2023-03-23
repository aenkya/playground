package playground

type Position int

type Snake struct {
	start int
	end   int
}

type Ladder struct {
	start int
	end   int
}

func addLaddersToBoard(board *Board) []*Ladder {
	ladders := []*Ladder{
		{start: 6, end: 25},
		{start: 11, end: 40},
		{start: 60, end: 85},
		{start: 46, end: 90},
	}
	board.ladderPositions = make(map[int]Position, len(ladders))

	for _, ladder := range ladders {
		board.ladderPositions[ladder.start] = Position(ladder.end)
	}
	return ladders
}

func addSnakesToBoard(board *Board) []*Snake {
	snakes := []*Snake{
		{start: 99, end: 54},
		{start: 70, end: 55},
		{start: 52, end: 42},
		{start: 25, end: 2},
	}
	board.snakePositions = make(map[int]Position, len(snakes))

	for _, snake := range snakes {
		board.snakePositions[snake.start] = Position(snake.end)
	}
	return snakes
}
