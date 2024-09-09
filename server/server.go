package main

import (
	"errors"
	"fmt"
	"net"
)

const (
	port = 4444
	ip   = "127.0.0.1"
)

var CurrentPlayers []net.Conn

func broadcast(data []byte) error {
	var errorList []error
	for _, player := range CurrentPlayers {
		if err := send(player, data); err != nil {
			errorList = append(errorList, err)
		}
	}
	return errors.Join(errorList...)
}

func send(conn net.Conn, data []byte) error {
	if _, err := conn.Write(data); err != nil {
		return err
	}
	return nil
}

func main() {
	println("Server is waiting...")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	for len(CurrentPlayers) < 2 {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		if err = send(conn, []byte("Hello from server")); err != nil {
			panic(err)
		}
		fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())
		CurrentPlayers = append(CurrentPlayers, conn)
	}

	println("Game can start now :)")
	if err = broadcast([]byte("Starting")); err != nil {
		panic(err)
	}
}
