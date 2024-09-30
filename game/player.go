package game

type Player struct {
	score float32
	Token Token
	Name  string
}

func NewPlayer(name string) *Player {
	return &Player{Name: name}
}
