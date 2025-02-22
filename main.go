package main

import (
	"encoding/hex"
	"fmt"
	"net"
)

func connectedToServer(host string, port int) error {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Println("Connected to server.")

	hexString := "000000230012674a4f74d28b00096b61666b612d636c69000a6b61666b612d636c6904302e3100"
	message, err := hex.DecodeString(hexString)
	if err != nil {
		return err
	}

	_, err = conn.Write(message)
	if err != nil {
		return err
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}

	fmt.Print("Server response: ", string(buffer[:n]))
	return nil
}

func main() {
	configuration, err := LoadConfiguration("config.json")
	if err != nil {
		return
	}

	connectedToServer(configuration.Host, configuration.Port)

	fmt.Println(configuration)
}
