package main

import "fmt"

func main() {
	board := NewBoard()
	gameState := NewGameState()
	for !gameState.GameFinished {
		token := gameState.CurrentToken()

		print(board.Representation())
		fmt.Printf("%c: Place your token\n", token)
		var location int
		if _, err := fmt.Scan(&location); err != nil {
			println("There was a problem with your input. Enter the number above the desired column")
			continue
		}

		row, err := board.Place(Token(token), location-1)
		if err != nil {
			println(err.Error())
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
