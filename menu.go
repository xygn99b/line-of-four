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
	fmt.Printf("<%c> Single player\n", Singleplayer)
	fmt.Printf("<%c> Multiplayer\n", Multiplayer)
	fmt.Printf("<%c> Exit\n", Exit)
	fmt.Print("\n>")

	var selection rune
	if _, err := fmt.Scanf("%c", &selection); err != nil {
		panic(err)
	}

	switch selection {
	case Singleplayer:
		game := NewGame(4)
		game.Run()
	case Multiplayer:
		os.Exit(0)
	case Exit:
		os.Exit(0)
	}
}
