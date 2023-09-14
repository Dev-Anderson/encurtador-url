package database

import (
	"database/sql"
	"encurtador-url/cmd/config"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectionDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.LoadEnv().Host, config.LoadEnv().User, config.LoadEnv().Password, config.LoadEnv().DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Falha ao fazer um ping no banco de dados", err.Error())
	}

	return db, nil

}
