package main

import "github.com/fatih/color"

type Token rune

const (
	TokenNull Token = iota // Blank spaces
	TokenRed        = 'r'
	TokenBlue       = 'b'
)

func (token Token) String() string {
	switch token {
	case TokenRed:
		return "Red"
	case TokenBlue:
		return "Blue"
	}
	return ""
}

func (token Token) Color() *color.Color {
	switch token {
	case TokenRed:
		return color.New(color.FgRed)
	case TokenBlue:
		return color.New(color.FgBlue)
	}
	return color.New(color.FgWhite)
}
