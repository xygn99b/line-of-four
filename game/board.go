package game

import "errors"

const (
	BoardColumns = 7
	BoardRows    = 6
)

var (
	ErrorColumnFull      = errors.New("The selected column is full")
	ErrorInvalidPosition = errors.New("The given position is invalid")
)

type Board struct {
	Locations                [BoardColumns][BoardRows]Token
	ConsecutiveWinningTokens int // Number of tokens in a line required to win
}

func NewBoard(consecutiveWinningTokens int) *Board {
	board := &Board{ConsecutiveWinningTokens: consecutiveWinningTokens}
	return board
}

// PrintRepresentation prints a string representation of the board and its tokens.
func (b *Board) PrintRepresentation() {
	print("  (1) (2) (3) (4) (5) (6) (7)\n")
	print("   -   -   -   -   -   -   -\n")
	for y := len(b.Locations[0]) - 1; y >= 0; y-- {
		for x := range len(b.Locations) {
			token := b.Locations[x][y]
			print("   ")
			if token == TokenNull {
				print("O")
				continue
			}
			token.Color().Print("O")
		}
		print("\n")
	}
	print("\n")
}

// Place places a token in the given column of the board, returning its row. It will return an ErrorColumnFull if the column already contains the maximum number of tokens
func (b *Board) Place(token Token, column int) (int, error) {
	if !b.ValidPosition(column, 0) {
		return -1, ErrorInvalidPosition
	}
	for i := range b.Locations[column] {
		if b.Locations[column][i] == TokenNull {
			b.Locations[column][i] = token
			return i, nil
		}
	}
	return -1, ErrorColumnFull
}

func (b *Board) maxConsecutiveFrom(location [2]int) int {
	token := b.At(location[0], location[1])
	if token == TokenNull {
		return 0
	}
	maxInARow := max(b.countRow(token, location[1]), b.countColumn(token, location[0]))

	// check diagonals
	var baseX, baseY int = 0, location[1] - location[0]
	if location[0] > location[1] {
		baseX, baseY = location[0]-location[1], 0
	}

	maxInARow = max(b.countForwardDiagonal(token, baseX, baseY), maxInARow)

	maxInARow = max(maxInARow, b.countBackwardDiagonal(token, baseX, BoardColumns-1-baseY))
	return maxInARow
}

// CheckWin returns isWin as true if the token at the location creates a winning scenario.
// If isWin, then locations will contain the locations of the winning tokens.
func (b *Board) CheckWin(location [2]int) (isWin bool) {
	return b.maxConsecutiveFrom(location) >= b.ConsecutiveWinningTokens
}

func (b *Board) countRow(token Token, row int) int {
	count := 0
	maxCount := 0
	for x := range BoardColumns {
		if b.At(x, row) != token {
			count = 0
			continue
		}
		count++
		maxCount = max(maxCount, count)
	}
	return maxCount
}

func (b *Board) countColumn(token Token, column int) int {
	count := 0
	maxCount := 0
	for y := range BoardRows {
		if b.At(column, y) != token {
			count = 0
			continue
		}
		count++
		maxCount = max(maxCount, count)
	}
	return maxCount
}

func (b *Board) countForwardDiagonal(token Token, baseX, baseY int) int {
	count := 0
	maxCount := 0
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
		maxCount = max(maxCount, count)
	}
	return maxCount
}

func (b *Board) countBackwardDiagonal(token Token, baseX, baseY int) int {
	count := 0
	maxCount := 0
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
		maxCount = max(maxCount, count)
	}
	return maxCount
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

// Full returns true if all locations on the board are occupied
func (b *Board) Full() bool {
	for _, row := range b.Locations {
		for _, token := range row {
			if token == TokenNull {
				return false
			}
		}
	}
	return true
}
