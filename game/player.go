package game

import (
	"lineof4/network"
	"net"
)

type Player struct {
	score float32
	Token Token
	conn  net.Conn
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) send(message network.Message) {
	p.conn.Write([]byte(message.Encode()))
}
