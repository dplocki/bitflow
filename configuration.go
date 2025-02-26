package main

import (
	"github.com/hjson/hjson-go/v4"
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

	err = hjson.Unmarshal(bytes, &configuration)
	if err != nil {
		return configuration, err
	}

	return configuration, nil
}
