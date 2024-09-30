package game

import (
	"fmt"
	"lineof4/utils"
	"net"
	"os"
)

const (
	Singleplayer      rune = 's'
	MultiplayerOnline rune = 'o'
	MultiplayerLocal  rune = 'l'
	Exit              rune = 'e'
)

func MainMenu() {
	utils.ClearScreen()
	println("Welcome to LINE OF X")
	println("Enter an option to play")
	println("----------------------")
	fmt.Printf("<%c> Single player (vs CPU)\n", Singleplayer)
	fmt.Printf("<%c> Local multiplayer\n", MultiplayerLocal)
	fmt.Printf("<%c> Online multiplayer\n", MultiplayerOnline)
	fmt.Printf("<%c> Exit\n", Exit)
	fmt.Print("\n>")

	var selection rune
	if _, err := fmt.Scanf("%c", &selection); err != nil {
		panic(err)
	}

	switch selection {
	case Singleplayer:
		os.Exit(0)
	case MultiplayerOnline:
		OnlineMenu()
	case MultiplayerLocal:
		game := NewGame(LocalGameType, 4)
		game.Run()
	case Exit:
		os.Exit(0)
	}
}

const (
	CreateGame rune = 'c'
	JoinGame   rune = 'j'
)

func OnlineMenu() {
	utils.ClearScreen()
	fmt.Printf("<%c> Create game\n", CreateGame)
	fmt.Printf("<%c> Join game\n", JoinGame)
	fmt.Printf("\n>")

	var selection rune
	if _, err := fmt.Scanf("\n%c", &selection); err != nil {
		panic(err)
	}

	switch selection {
	case CreateGame:
		os.Exit(0)
	case JoinGame:
		fmt.Printf("Enter IP:PORT [example: 127.0.0.1:4444]\n>")
		var address string
		if _, err := fmt.Scanf("\n%s", &address); err != nil {
			panic(err)
		}
		conn, err := net.Dial("tcp", address)
		if err != nil {
			panic(err)
		}

		for {
			var buffer []byte = make([]byte, 128)
			_, err = conn.Read(buffer)
			if err != nil {
				panic(err)
			}
			println(string(buffer))

		}
	}
}
