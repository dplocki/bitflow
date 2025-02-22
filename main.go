package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"io/ioutil"
	"os"
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

type Person struct {
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	Email    string   `json:"email"`
	Messages []string `json:"messages"`
}

func main() {
	// Open JSON file
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read file content
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse JSON into a generic map
	var data map[string]interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print the data
	fmt.Println("Loaded JSON:", data)
}
