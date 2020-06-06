package main

import (
	"errors"
	"fmt"
)

const (
	PlayerX int = 1
	PlayerO int = -1

	StateWonX = 2
	StateWonO = 3
	StateDraw = 4
	StateOpen = 5
)

type Game struct {
	Board [3][3]int
}

func (g *Game) String() string {
	s := ""

	spot := 1
	for x, row := range g.Board {
		for y, _ := range row {
			switch g.Board[x][y] {
			case PlayerX:
				s += fmt.Sprintf("X ")
			case PlayerO:
				s += fmt.Sprintf("O ")
			default:
				s += fmt.Sprintf("%d ", spot)
			}
			spot += 1
		}
		s += fmt.Sprintf("\n")
	}

	return s
}

func (g *Game) Play(player int, spot int) error {
	i := 1
	for x := range g.Board {
		for y := range g.Board[x] {
			if i == spot {
				if g.Board[x][y] == 1 || g.Board[x][y] == 2 {
					return errors.New("Spot taken!")
				} else {
					g.Board[x][y] = player

					return nil
				}
			}
			i++
		}
	}

	return nil
}

func (g *Game) State() int {
	if g.Won(PlayerO) {
		return StateWonO
	}

	if g.Won(PlayerX) {
		return StateWonX
	}

	for x, row := range g.Board {
		for y, _ := range row {
			if g.Board[x][y] != PlayerX && g.Board[x][y] != PlayerO {
				return StateOpen
			}
		}
	}

	return StateDraw
}

func (g *Game) Won(player int) bool {
	// Check rows
	for row := range g.Board {
		i := 0
		for col := range g.Board[row] {
			if g.Board[row][col] == player {
				i++
			}
			if i == 3 {
				return true
			}
		}
	}

	// Check columns
	for row := range g.Board {
		i := 0
		for col := range g.Board[row] {
			if g.Board[col][row] == player {
				i++
			}
			if i == 3 {
				return true
			}
		}
	}

	// Check diagonals
	if g.Board[0][0] == player &&
		g.Board[1][1] == player &&
		g.Board[2][2] == player {
		return true
	}
	if g.Board[0][2] == player &&
		g.Board[1][1] == player &&
		g.Board[2][0] == player {
		return true
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func (g *Game) MiniMax(depth int, isMax bool) int {
	switch g.State() {
	case StateWonO:
		return 10
	case StateWonX:
		return -10
	case StateDraw:
		return 0
	}

	best, player := -999, PlayerO
	if !isMax {
		best, player = 999, PlayerX
	}

	for x := range g.Board {
		for y := range g.Board[x] {
			if g.Board[x][y] == 0 {
				g.Board[x][y] = player
				if isMax {
					best = max(best, g.MiniMax(depth+1, !isMax))
				} else {
					best = min(best, g.MiniMax(depth+1, !isMax))
				}
				g.Board[x][y] = 0
			}
		}
	}

	return best
}

func (g *Game) BestMove() int {
	bestScore := -999
	bestMove := 0

	i := 1
	for x := range g.Board {
		for y := range g.Board[x] {
			if g.Board[x][y] == 0 {
				g.Board[x][y] = PlayerO
				currentScore := g.MiniMax(0, false)

				g.Board[x][y] = 0

				if currentScore > bestScore {
					bestScore = currentScore
					bestMove = i
				}
			}
			i++
		}
	}

	return bestMove
}

func main() {
	g := new(Game)

	for g.State() == StateOpen {
		fmt.Println(g)
		var spot int
		fmt.Printf("Spot: ")
		fmt.Scan(&spot)
		fmt.Println()

		err := g.Play(PlayerX, spot)
		if err != nil {
			fmt.Println(err)
			fmt.Println()
		}

		if g.State() == StateWonX {
			fmt.Println(g)
			fmt.Println("X won!")
			return
		}

		g.Play(PlayerO, g.BestMove())
		if g.State() == StateWonO {
			fmt.Println(g)
			fmt.Println("O won!")
			return
		}

		if g.State() == StateDraw {
			fmt.Println("Draw!")
		}
	}
}
