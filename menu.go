package main

import (
	"fmt"
	"os"
)

func MainMenu() {
	ClearScreen()
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
		game := NewGame(4)
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
	ClearScreen()
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
		os.Exit(0)
	}
}
