package main

import "github.com/fatih/color"

type Token rune

const (
	TokenNull Token = iota // Blank spaces
	TokenRed        = 'r'
	TokenBlue       = 'b'
)

func GetTokenString(token Token) string {
	switch token {
	case TokenRed:
		return "Red"
	case TokenBlue:
		return "Blue"
	}
	return ""
}

func GetTokenColor(token Token) *color.Color {
	switch token {
	case TokenRed:
		return color.New(color.BgRed)
	case TokenBlue:
		return color.New(color.BgBlue)
	}
	return color.New(color.FgWhite)
}
