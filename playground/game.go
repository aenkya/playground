package playground

import (
	"fmt"
)

type Game interface {
	Start()
}

type SnakesAndLaddersGame struct {
	board   *Board
	snakes  []*Snake
	ladders []*Ladder
	players []*Player
}

func (s *SnakesAndLaddersGame) Start() {
	fmt.Println("Welcome to Snakes & Ladders!")

	s.getGameDetails()
	s.runGameEngine()
}

/*
runGameEngine runs the game and only stops when a player wins
*/
func (s *SnakesAndLaddersGame) runGameEngine() {
	for {
		for _, player := range s.players {
			fmt.Printf("\n%s's turn. Press enter to roll the dice.", player.name)
			fmt.Scanln()

			dice := player.RollDice()
			fmt.Printf("\n%s rolled a %d", player.name, dice)

			s.board.MovePlayer(player, Position(dice))

			if s.board.CheckWinner(player) {
				fmt.Printf("\n%s won the game!", player.name)
				return
			}
		}
	}
}

/*
getGameDetails gets the details of the game from the users
*/
func (s *SnakesAndLaddersGame) getGameDetails() {
	var numOfPlayers int

	fmt.Print("Enter number of players: ")

	_, err := fmt.Scanln(&numOfPlayers)
	if err != nil {
		fmt.Println("Invalid number of players.")
	}

	for i := 0; i < numOfPlayers; i++ {
		name := ""

		fmt.Printf("\nEnter name of player %d: ", i+1)
		fmt.Scan(&name)

		token := ""

		fmt.Printf("\nEnter token of player %d: ", i+1)
		fmt.Scan(&token)
		s.players = append(s.players, &Player{name: name, token: token})
	}
}

/*
NewGame initializes the game with the board, snakes, ladders and players
*/
func NewGame() Game {
	boardSize := 100
	board := &Board{size: boardSize}
	snakes := addSnakesToBoard(board)
	ladders := addLaddersToBoard(board)
	game := &SnakesAndLaddersGame{
		board:   board,
		snakes:  snakes,
		ladders: ladders,
	}

	return game
}
