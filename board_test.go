package main

import "testing"

func init() {
}

func TestBoardPlace(t *testing.T) {
	board := NewBoard()
	if _, err := board.Place(TokenRed, 4); err != nil {
		t.Error(err)
	}
	if board.Locations[4][0] != TokenRed {
		t.Fatal("Board does not place tokens correctly")
	}
}

func TestBoardPlaceFull(t *testing.T) {
	board := NewBoard()
	for range 6 {
		if _, err := board.Place(TokenBlue, 0); err != nil {
			t.Error(err)
		}
	}
	if _, err := board.Place(TokenBlue, 0); err != ErrorColumnFull {
		t.Fatal("Board does not prevent placing into a full column")
	}
}

func TestBoardCheckWinEmpty(t *testing.T) {
	board := NewBoard()
	if win := board.CheckWin([2]int{0, 0}); win {
		t.Fatal("Board considers an empty board a winning scenario")
	}
}

func TestBoardCheckWinRow(t *testing.T) {
	board := NewBoard()
	board.Place(TokenBlue, 2)
	board.Place(TokenBlue, 3)
	board.Place(TokenBlue, 4)
	row, _ := board.Place(TokenBlue, 5)
	if win := board.CheckWin([2]int{5, row}); !win {
		t.Fatal("Board does not acknowledge row wins")
	}
}

func TestBoardCheckWinColumn(t *testing.T) {
	board := NewBoard()
	board.Place(TokenBlue, 2)
	board.Place(TokenBlue, 2)
	board.Place(TokenBlue, 2)
	row, _ := board.Place(TokenBlue, 2)
	if win := board.CheckWin([2]int{2, row}); !win {
		t.Fatal("Board does not acknowledge column wins")
	}
}

func TestBoardCheckWinForwardDiagonal(t *testing.T) {
	board := NewBoard()

	for i := range 4 {
		for range i + 1 {
			board.Place(TokenRed, i)
		}
	}

	if win := board.CheckWin([2]int{1, 1}); !win {
		t.Fatal("Board does not acknowledge forward diagonal wins")
	}
}

func TestBoardCheckWinBackwardDiagonal(t *testing.T) {
	board := NewBoard()

	for i := range 4 {
		for range i + 1 {
			board.Place(TokenRed, 5-i)
		}
	}

	if win := board.CheckWin([2]int{5, 0}); !win {
		t.Fatal("Board does not acknowledge backward diagonal wins")
	}
}

func TestBoardFullEmpty(t *testing.T) {
	board := NewBoard()

	if board.Full() {
		t.Fatal("Board considers an empty grid 'full'")
	}
}

func TestBoardFull(t *testing.T) {
	board := NewBoard()

	for i := range 7 {
		for range 6 {
			board.Place(TokenBlue, i)
		}
	}

	if !board.Full() {
		t.Fatal("A full board does not consider itself 'full'.")
	}
}
