package game

import (
	"fmt"
	"lineof4/network"
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
		game, err := NewGame(LocalGameType, 4)
		if err != nil {
			panic(err)
		}
		players := []*Player{NewPlayer(), NewPlayer()}
		if err = game.Run(players); err != nil {
			panic(err)
		}
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
		addr := network.Launch(true)
		fmt.Printf("Server launched at %s\n", addr.String())
		joinGame(addr.String())
	case JoinGame:
		fmt.Printf("Enter IP:PORT [example: 127.0.0.1:4444]\n>")
		var address string
		if _, err := fmt.Scanf("\n%s", &address); err != nil {
			panic(err)
		}
		joinGame(address)
	}
}

func joinGame(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	conn.Write([]byte(network.NewMessage(network.WelcomeMessage, "hi").Encode()))
	var pos rune = 0
	for {
		buf := make([]byte, 128)
		_, err = conn.Read(buf)
		if err != nil {
			panic(err)
		}
		msg, err := network.NewMessageFromBytes(buf)
		if err != nil {
			println(err.Error())
			continue
		}
		if msg.Type == network.StartMessage {
			break
		}
		if msg.Type == network.WelcomeMessage {
			pos = rune(msg.Payload[0])
		}
	}
	game, err := NewGame(OnlineGameType, 4)
	focusPlayer := NewPlayer()
	players := []*Player{NewPlayer(), focusPlayer}
	if pos == rune(0) {
		players = []*Player{focusPlayer, NewPlayer()}
	}
	game.FocusPlayer = focusPlayer
	game.Conn = conn
	println("Starting")
	if err = game.Run(players); err != nil {
		panic(err)
	}
}
