package main

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	Email    string   `json:"email"`
	Messages []string `json:"messages"`
}

func LoadConfiguration(filename string) (Configuration, error) {
	var configuration Configuration

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return configuration, err
	}

	err = json.Unmarshal(bytes, &configuration)
	if err != nil {
		return configuration, err
	}

	return configuration, nil
}
