package main

import "fmt"

func main() {
	board := NewBoard()
	gameState := NewGameState()
	for !gameState.GameFinished {
		token := gameState.CurrentToken()

		print(board.Representation())
		fmt.Printf("%c: Place your token", token)
		var location int
		if _, err := fmt.Scan(&location); err != nil {
			panic(err)
		}

		row, err := board.Place(Token(token), location-1)
		if err == ErrorColumnFull {
			println("The selected column is full.")
			continue
		}

		gameState.NextTurn()

		if board.Full() {
			gameState.GameFinished = true
			gameState.EndGameMessage = "Board is full. Draw"
		}

		win := board.CheckWin([2]int{location - 1, row})
		if win {
			gameState.GameFinished = true
			gameState.EndGameMessage = fmt.Sprintf("%c won!", token)
		}

	}
	print(board.Representation())
	println(gameState.EndGameMessage)
}
