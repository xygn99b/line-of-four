package main

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
