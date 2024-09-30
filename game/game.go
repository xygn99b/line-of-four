package game

import (
	"errors"
	"fmt"
	"lineof4/utils"
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
	Players        []*Player
	GameFinished   bool
	EndGameMessage string
}

var ErrorTooManyPlayers = errors.New("too many players for the game")

func NewGameState(playerList []*Player) (*GameState, error) {
	gameState := &GameState{}
	gameState.Tokens = []Token{TokenBlue, TokenRed}
	gameState.Players = playerList
	if len(gameState.Players) > len(gameState.Tokens) {
		return nil, ErrorTooManyPlayers
	}
	for i, player := range gameState.Players {
		player.Token = gameState.Tokens[i]
	}
	return gameState, nil
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

func NewGame(gameType GameType, consecutiveWinningTokens int) (Game, error) {
	g := Game{GameType: gameType, consecutiveWinningTokens: consecutiveWinningTokens}
	return g, nil
}

func (g *Game) GetNextTokenPlaceLocation() (int, error) {
	switch g.GameType {
	case LocalGameType:
		return g.promptUserTokenPlace()
	}
	panic(errors.New("invalid game type"))
}

func (g *Game) promptUserTokenPlace() (int, error) {
	token := g.State.CurrentToken()
	token.Color().Printf("%s: Place your token", token)
	print("\n>")

	var location int
	if _, err := fmt.Scan(&location); err != nil {
		println("There was a problem with your input. Enter the number above the desired column")
		return -1, err
	}

	return location - 1, nil
}

func (g *Game) init(playerList []*Player) error {
	g.Board = NewBoard(g.consecutiveWinningTokens)
	state, err := NewGameState(playerList)
	if err != nil {
		return err
	}
	g.State = state
	return nil
}

func (g *Game) Run(playerList []*Player) error {
	if err := g.init(playerList); err != nil {
		return err
	}
	for !g.State.GameFinished {
		utils.ClearScreen()

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

	utils.ClearScreen()
	g.Board.PrintRepresentation()
	println(g.State.EndGameMessage)
	return nil
}
