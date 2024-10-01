package game

import (
	"errors"
	"fmt"
	"lineof4/network"
	"lineof4/utils"
	"net"
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
	Players        map[Token]*Player
	GameFinished   bool
	EndGameMessage string
}

var ErrorTooManyPlayers = errors.New("too many players for the game")

func NewGameState(playerList []*Player) (*GameState, error) {
	gameState := &GameState{}
	gameState.Tokens = []Token{TokenBlue, TokenRed}

	if len(gameState.Players) > len(gameState.Tokens) {
		return nil, ErrorTooManyPlayers
	}
	gameState.Players = make(map[Token]*Player)
	for i, token := range gameState.Tokens {
		player := playerList[i]
		gameState.Players[token] = player
		player.Token = token
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
	Message                  string
	FocusPlayer              *Player
	Conn                     net.Conn
}

func NewGame(gameType GameType, consecutiveWinningTokens int) (Game, error) {
	g := Game{GameType: gameType, consecutiveWinningTokens: consecutiveWinningTokens}
	return g, nil
}

func (g *Game) GetNextTokenPlaceLocation() (int, error) {
	switch g.GameType {
	case LocalGameType:
		return g.promptUserTokenPlace()
	case OnlineGameType:
		if g.isCurrentPlayerFocused() {
			return g.promptUserTokenPlace()
		}
		return g.getOnlineOpponentTokenPlace()
	}
	panic(errors.New("invalid game type"))
}
func (g *Game) isCurrentPlayerFocused() bool {
	return g.State.Players[g.State.CurrentToken()] == g.FocusPlayer
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

	if g.isCurrentPlayerFocused() {
		_, err := g.Conn.Write([]byte(network.NewMessage(network.PlaceMessage, string(byte(location-1))).Encode()))
		if err != nil {
			panic(err)
		}
	}

	return location - 1, nil
}
func (g *Game) getOnlineOpponentTokenPlace() (int, error) {
	println("Waiting for opponent...")
	if g.Conn == nil {
		return 0, errors.New("no connection was found")
	}
	for {
		buf := make([]byte, 128)
		if _, err := g.Conn.Read(buf); err != nil {
			panic(err)
		}
		msg, err := network.NewMessageFromBytes(buf)
		if err != nil {
			continue
		}
		if msg.Type == network.PlaceMessage {
			x := int(msg.Payload[0])
			return x, nil
		}
	}

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
		println(g.Message)
		token := g.State.CurrentToken()
		g.Board.PrintRepresentation()

		location, err := g.GetNextTokenPlaceLocation()

		if err != nil {
			g.Message = err.Error()
			continue
		}

		row, err := g.Board.Place(token, location)
		if err != nil {
			g.Message = err.Error()
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

		g.Message = ""
		g.State.NextTurn()
	}

	utils.ClearScreen()
	g.Board.PrintRepresentation()
	println(g.State.EndGameMessage)
	return nil
}
