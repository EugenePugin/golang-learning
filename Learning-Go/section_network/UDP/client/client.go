package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	serverAddress := "127.0.0.1:8081"

	// 1. Connect to the TCP server
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Printf("Error connecting to %s: %s\n", serverAddress, err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Printf("Connected to TCP server at %s\n", serverAddress)

	// 2. Read the response from the server using conn.Read
	// Define a buffer to hold the incoming data
	buffer := make([]byte, 1024) // 1024 bytes (1 KB) buffer size

	const NumOfMsgsToRead = 3
	for i := 0; i < NumOfMsgsToRead; i++ {
		// conn.Read reads data into the buffer
		// n is the number of bytes read
		// err indicates if an error occurred (e.g., io.EOF for end of file)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
			return
		}

		// Convert the read bytes (up to n) to a string
		receivedMessage := strings.ToUpper(string(buffer[:n]))

		fmt.Printf( /*"Received (%d bytes): %s", n, */ receivedMessage)
	}

}
