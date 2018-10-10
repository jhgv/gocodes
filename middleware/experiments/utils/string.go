package utils

import (
	"math/rand"
	"time"
	"strings"
)

func GenerateRandomText(maxLength int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	randomText := make([]rune, maxLength)
	for i := range randomText {
		randomText[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(randomText)
}

func ProcessedMessage(message string) string {
	return strings.ToUpper(strings.TrimRight(message, "\n"))
}