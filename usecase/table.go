package usecase

import (
	"encurtador-url/database"
	"log"
)

const (
	createTable = `
		CREATE TABLE IF NOT EXISTS encurtador_url (
			id SERIAL PRIMARY KEY, 
			codigo VARCHAR(255), 
			url_original VARCHAR(255), 
			data_create TIMESTAMPTZ DEFAULT NOW(), 
			data_update TIMESTAMPTZ, 
			data_delete TIMESTAMPTZ
		)
	`
)

func CreateTable() error {
	db, err := database.ConnectDatabse()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
