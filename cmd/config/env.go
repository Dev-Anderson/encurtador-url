package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigDB struct {
	User       string
	Password   string
	Host       string
	Port       string
	DBName     string
	PortServer string
	LengthUrl  string
	Caracteres string
}

func LoadEnv() ConfigDB {
	godotenv.Load()
	return ConfigDB{
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		PortServer: os.Getenv("PORT_SERVER"),
		LengthUrl:  os.Getenv("LENGTH_URL"),
		Caracteres: os.Getenv("RANDOM_CARACTERES"),
	}
}
