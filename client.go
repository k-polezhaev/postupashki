package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		os.Exit(1)
	}
	defer func() {
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')

	if err != nil {
		return
	}

	if response == "OK" {
		fmt.Println("Response 'OK' successfully received and verified.")
	} else {
		fmt.Printf("Error: Unexpected response received: \"%s\"\n", response)
	}
}
