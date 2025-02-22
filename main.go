package main

import (
	"encoding/hex"
	"fmt"
	"net"
)

func connectedToServer() {
	conn, err := net.Dial("tcp", "localhost:9092")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server.")

	hexString := "000000230012674a4f74d28b00096b61666b612d636c69000a6b61666b612d636c6904302e3100"
	message, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return
	}

	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Print("Server response: ", string(buffer[:n]))
}

func main() {
	configuration, err := LoadConfiguration("config.json")
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	fmt.Println(configuration)
}