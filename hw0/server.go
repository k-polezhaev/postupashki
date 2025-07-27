package main

import (
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	_, err := conn.Write([]byte("OK\n"))
	if err != nil {
		return
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}
