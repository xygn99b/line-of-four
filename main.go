package main

import "fmt"

func main() {
	board := NewBoard()
	turnNumber := 0
	turnOver := false
	token := TokenRed
	for {
		if turnOver {
			turnNumber++
			if turnNumber%2 == 0 {
				token = TokenRed
			} else {
				token = TokenBlue
			}
			turnOver = false
		}

		print(board.Representation())
		fmt.Printf("%c: Place your token", token)
		var location int
		if _, err := fmt.Scan(&location); err != nil {
			panic(err)
		}

		row, err := board.Place(Token(token), location-1)
		if err == ErrorColumnFull {
			println("The selected column is full.")
			continue
		}

		turnOver = true

		win := board.CheckWin([2]int{location - 1, row})
		if win {
			print(board.Representation())
			fmt.Printf("%c wins\n", token)
			break
		}
	}
}
