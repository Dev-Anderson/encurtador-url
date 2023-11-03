package config

import (
	"bufio"
	"encurtador-url/schemas"
	"os"
	"strings"
)

func setEnv(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func LoadEnv() schemas.Env {
	setEnv(".env")
	return schemas.Env{
		Host:             os.Getenv("DB_HOST"),
		Port:             os.Getenv("DB_PORT"),
		User:             os.Getenv("DB_USER"),
		Password:         os.Getenv("DB_PASSWORD"),
		DBName:           os.Getenv("DB_NAME"),
		PortServer:       os.Getenv("PORT_SERVER"),
		RandomCaracteres: os.Getenv("RANDOM_CARACTERES"),
	}
}
