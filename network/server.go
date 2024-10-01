package network

import (
	"errors"
	"fmt"
	"net"
)

const (
	defaultPort = 0
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

var CurrentPlayers []net.Conn

func broadcast(data []byte, exclude net.Conn) error {
	var errorList []error
	for _, player := range CurrentPlayers {
		if player == exclude {
			continue
		}
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

func Launch(silent bool) net.Addr {
	if !silent {
		println("Server is waiting...")
	}
	ip := GetLocalIP()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, defaultPort))
	if err != nil {
		panic(err)
	}
	go run(listener, silent)
	return listener.Addr()
}

func run(listener net.Listener, silent bool) {
	for len(CurrentPlayers) < 2 {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		if err = send(conn, []byte(NewMessage(WelcomeMessage, string(rune(len(CurrentPlayers)))).Encode())); err != nil {
			panic(err)
		}
		if !silent {
			fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())
		}
		CurrentPlayers = append(CurrentPlayers, conn)
	}
	if !silent {
		println("Game is starting")
	}
	if err := broadcast([]byte(NewMessage(StartMessage, "").Encode()), nil); err != nil {
		panic(err)
	}
	for _, player := range CurrentPlayers {
		go handle(player)
	}
}
func handle(conn net.Conn) {
	for {
		buf := make([]byte, 128)
		_, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		if err := broadcast(buf, conn); err != nil {
			panic(err)
		}
	}
}
