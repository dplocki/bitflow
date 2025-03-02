package main

import (
	"encoding/hex"
)

func ParsingMessages(messages []string) ([][]byte, error) {
    result := make([][]byte, len(messages))
    for index, message := range messages {
        bytes, err := hex.DecodeString(message)
        if err != nil {
            return nil, err
        }
        result[index] = bytes
    }
    return result, nil
}
