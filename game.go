package main

import (
	"fmt"
)

type GameState struct {
	TurnCount      int
	Tokens         []Token
	GameFinished   bool
	EndGameMessage string
}

func NewGameState() *GameState {
	gameState := &GameState{}
	gameState.Tokens = []Token{TokenBlue, TokenRed}
	return gameState
}

func (gs *GameState) NextTurn() {
	gs.TurnCount++
}

// CurrentToken returns the token which is next to be placed on this turn
func (gs *GameState) CurrentToken() Token {
	return gs.Tokens[gs.TurnCount%len(gs.Tokens)]
}

type Game struct {
	State *GameState
}

func NewGame() Game {
	g := Game{}
	g.State = NewGameState()
	return g
}

func (g Game) Run() {
	board := NewBoard()
	gameState := NewGameState()
	for !gameState.GameFinished {
		token := gameState.CurrentToken()

		print(board.Representation())
		GetTokenColor(token).Printf("%s: Place your token", GetTokenString(token))
		print("\n>")

		var location int
		if _, err := fmt.Scan(&location); err != nil {
			println("There was a problem with your input. Enter the number above the desired column")
			continue
		}

		row, err := board.Place(token, location-1)
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
			gameState.EndGameMessage = GetTokenColor(token).Sprintf("%s won!", GetTokenString(token))
		}
	}
	print(board.Representation())
	println(gameState.EndGameMessage)
}
