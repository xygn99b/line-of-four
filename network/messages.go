package network

import (
	"errors"
	"strings"
)

type MessageType string

const MessageDelimiter = '|'

const (
	StartMessage   MessageType = "s"
	PlaceMessage   MessageType = "p"
	WelcomeMessage MessageType = "w"
)

type Message struct {
	Type    MessageType
	Payload string
}

func NewMessage(messageType MessageType, payload string) Message {
	return Message{messageType, payload}
}
func NewMessageFromBytes(in []byte) (Message, error) {
	parts := strings.Split(string(in), string(MessageDelimiter))
	if len(parts) < 2 {
		return Message{}, errors.New("invalid message")
	}
	return Message{MessageType(parts[0]), parts[1]}, nil
}
func (m Message) Encode() string {
	out := strings.Builder{}
	out.WriteString(string(m.Type))
	out.WriteRune(MessageDelimiter)
	out.WriteString(m.Payload)
	return out.String()
}
