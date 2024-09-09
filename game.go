package main

import (
	"errors"
	"fmt"
)

type GameType int

const (
	OnlineGameType = iota
	CPUGameType
	LocalGameType
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
	State                    *GameState
	Board                    *Board
	GameType                 GameType
	consecutiveWinningTokens int
}

func NewGame(gameType GameType, consecutiveWinningTokens int) Game {
	g := Game{GameType: gameType, consecutiveWinningTokens: consecutiveWinningTokens}
	g.State = NewGameState()
	return g
}

func (g Game) GetNextTokenPlaceLocation() (int, error) {
	switch g.GameType {
	case LocalGameType:
		return g.promptUserTokenPlace()
	}
	panic(errors.New("invalid game type"))
}

func (g Game) promptUserTokenPlace() (int, error) {
	token := g.State.CurrentToken()
	token.Color().Printf("%s: Place your token", token)
	print("\n>")

	var location int
	if _, err := fmt.Scan(&location); err != nil {
		println("There was a problem with your input. Enter the number above the desired column")
		return -1, err
	}

	return location, nil
}

func (g Game) Run() {
	g.Board = NewBoard(g.consecutiveWinningTokens)
	g.State = NewGameState()
	for !g.State.GameFinished {
		ClearScreen()

		token := g.State.CurrentToken()
		g.Board.PrintRepresentation()

		location, err := g.GetNextTokenPlaceLocation()
		if err != nil {
			continue
		}

		row, err := g.Board.Place(token, location)
		if err != nil {
			println(err.Error())
			continue
		}

		win := g.Board.CheckWin([2]int{location, row})
		if win {
			g.State.GameFinished = true
			g.State.EndGameMessage = token.Color().Sprintf("%s won!", token)
		}

		if g.Board.Full() {
			g.State.GameFinished = true
			g.State.EndGameMessage = "Board is full. Draw"
		}

		g.State.NextTurn()
	}

	ClearScreen()
	g.Board.PrintRepresentation()
	println(g.State.EndGameMessage)
}
