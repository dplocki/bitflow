package main

import (
	"encoding/hex"
    "strings"
)

func normalizeMessage(message string) string {
    if strings.Contains(message, "|") {
        parts := strings.Split(message, "|")
        if len(parts) > 1 {
            message = parts[1]
        }
    }

    return strings.ReplaceAll(message, " ", "")
}

func ParsingMessages(messages []string) ([][]byte, error) {
    result := make([][]byte, len(messages))

    for index, message := range messages {
        message = normalizeMessage(message)
        bytes, err := hex.DecodeString(message)
        if err != nil {
            return nil, err
        }
        result[index] = bytes
    }

    return result, nil
}
