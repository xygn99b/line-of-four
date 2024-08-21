package main

import "errors"

type Token rune

const (
	BoardColumns             = 7
	BoardRows                = 6
	ConsecutiveWinningTokens = 4 // Number of tokens in a row required to win the game (connect X)
)

const (
	TokenNull Token = iota // Blank spaces
	TokenRed        = 'r'
	TokenBlue       = 'b'
)

var ErrorColumnFull = errors.New("the selected column is full")

type Board struct {
	Locations [BoardColumns][BoardRows]Token
}

func NewBoard() *Board {
	board := &Board{}
	return board
}

// Representation returns a string representation of the board and its tokens.
func (b *Board) Representation() string {
	var out string
	out += "  (1) (2) (3) (4) (5) (6) (7)\n"
	out += "   -   -   -   -   -   -   -\n"
	for y := len(b.Locations[0]) - 1; y >= 0; y-- {
		var rowString string
		for x := range len(b.Locations) {
			token := b.Locations[x][y]
			rowString += "   "
			if token == TokenNull {
				rowString += "O"
				continue
			}
			rowString += string(token)
		}
		out += rowString + "\n"
	}
	return out
}

// Place places a token in the given column of the board, returning its row. It will return an ErrorColumnFull if the column already contains the maximum number of tokens
func (b *Board) Place(token Token, column int) (int, error) {
	for i := range b.Locations[column] {
		if b.Locations[column][i] == TokenNull {
			b.Locations[column][i] = token
			return i, nil
		}
	}
	return -1, ErrorColumnFull
}

// CheckWin returns isWin as true if the token at the location creates a winning scenario.
// If isWin, then locations will contain the locations of the winning tokens.
func (b *Board) CheckWin(location [2]int) (isWin bool) {
	token := b.At(location[0], location[1])
	if token == TokenNull {
		return false
	}

	if b.checkWinRow(token, location[1]) {
		return true
	}

	if b.checkWinColumn(token, location[0]) {
		return true
	}

	// check diagonals
	var baseX, baseY int = 0, location[1] - location[0]
	if location[0] > location[1] {
		baseX, baseY = location[0]-location[1], 0
	}

	if b.checkWinForwardDiagonal(token, baseX, baseY) {
		return true
	}

	if b.checkWinBackwardDiagonal(token, baseX, BoardColumns-1-baseY) {
		return true
	}

	return false
}

func (b *Board) checkWinRow(token Token, row int) bool {
	var count int = 0
	for x := range BoardColumns {
		if b.At(x, row) != token {
			count = 0
			continue
		}
		count++
		if count == ConsecutiveWinningTokens {
			return true
		}
	}
	return false
}

func (b *Board) checkWinColumn(token Token, column int) bool {
	count := 0
	for y := range BoardRows {
		if b.At(column, y) != token {
			count = 0
			continue
		}
		count++
		if count == ConsecutiveWinningTokens {
			return true
		}
	}
	return false
}

func (b *Board) checkWinForwardDiagonal(token Token, baseX, baseY int) bool {
	count := 0

	for i := range max(BoardRows, BoardColumns) {
		x, y := baseX+i, baseY+i
		if !b.ValidPosition(x, y) {
			break
		}

		if token != b.At(x, y) {
			count = 0
			continue
		}
		count++
		if count == ConsecutiveWinningTokens {
			return true
		}
	}

	return false
}

func (b *Board) checkWinBackwardDiagonal(token Token, baseX, baseY int) bool {
	count := 0

	for i := range max(BoardColumns, BoardRows) {
		x, y := baseX+i, baseY-i
		if !b.ValidPosition(x, y) {
			break
		}

		if token != b.At(x, y) {
			count = 0
			continue
		}
		count++
		if count == ConsecutiveWinningTokens {
			return true
		}
	}
	return false
}

func (b *Board) ValidPosition(x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x > BoardColumns-1 || y > BoardRows-1 {
		return false
	}
	return true
}

// At returns the token at the given location
func (b *Board) At(x, y int) Token {
	return b.Locations[x][y]
}
