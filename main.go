package main

import (
	"encoding/hex"
	"fmt"
	"net"
)

func connectedToServer(host string, port int, messages []string) error {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Print("Connected to server ")
	fmt.Print(host)
	fmt.Print(":")
	fmt.Println(port)

	for index, message := range messages {
		bytes, err := hex.DecodeString(message)
		if err != nil {
			return err
		}

		fmt.Print(index)
		hex.Dump(bytes)

		_, err = conn.Write(bytes)
		if err != nil {
			return err
		}

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Println("Server response:")
		hex.Dump(buffer[:n])
	}

	return nil
}

func main() {
	configuration, err := LoadConfiguration("config.json")
	if err != nil {
		fmt.Errorf("Error during the loading configuration: %w", err)
		return
	}

	connectedToServer(configuration.Host, configuration.Port, configuration.Messages)
	if err != nil {
		fmt.Errorf("Error during the server connection: %w", err)
		return
	}
}
