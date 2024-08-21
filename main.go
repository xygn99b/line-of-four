package main

import "fmt"

func main() {
	board := NewBoard()
	turn := 0
	token := TokenRed
	for {
		print(board.Representation())
		println("Place your token")
		var location int
		if _, err := fmt.Scan(&location); err != nil {
			panic(err)
		}
		turn++
		turn %= 2
		if turn%2 == 0 {
			token = TokenRed
		} else {
			token = TokenBlue
		}
		row, err := board.Place(Token(token), location-1)
		if err == ErrorColumnFull {
			println("The selected column is full.")
			continue
		}
		win := board.CheckWin([2]int{location - 1, row})
		if win {
			print(board.Representation())
			fmt.Printf("%c wins\n", token)
			break
		}
	}
}
