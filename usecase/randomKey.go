package usecase

import (
	"encurtador-url/config"
	"math/rand"
	"strconv"
	"time"
)

func Random() (string, error) {
	c := config.LoadEnv().RandomCaracteres
	charSetLength := len(config.LoadEnv().RandomLength)
	length, _ := strconv.Atoi(config.LoadEnv().RandomLength)

	rand.Seed(time.Now().UnixNano())
	usedChars := make(map[byte]bool)
	result := ""

	for i := 0; i < length; i++ {
		randIndex := rand.Intn(charSetLength)
		char := c[randIndex]

		for usedChars[char] {
			randIndex = rand.Intn(charSetLength)
			char = c[randIndex]
		}

		result += string(char)
		usedChars[char] = true
	}

	return result, nil
}
