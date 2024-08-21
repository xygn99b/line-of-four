package main

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
