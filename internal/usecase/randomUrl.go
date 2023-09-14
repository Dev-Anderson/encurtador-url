package usecase

import (
	"encurtador-url/cmd/config"
	"math/rand"
	"time"
)

func RandomUrl(length int) (string, error) {
	characters := config.LoadEnv().Caracteres
	charSetLength := len(characters)

	rand.Seed(time.Now().UnixNano())
	usedChars := make(map[byte]bool)
	result := ""

	for i := 0; i < length; i++ {
		randIndex := rand.Intn(charSetLength)
		char := characters[randIndex]

		for usedChars[char] {
			randIndex = rand.Intn(charSetLength)
			char = characters[randIndex]
		}

		result += string(char)
		usedChars[char] = true
	}

	return result, nil

}
