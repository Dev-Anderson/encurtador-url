package models

import (
	"encurtador-url/database"
	"encurtador-url/schemas"
	"encurtador-url/usecase"
	"fmt"
	"log"
)

func GetAllUrls() ([]schemas.URL, error) {
	db, err := database.ConnectDatabse()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "select codigo, url_original  from encurtador_url"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error get URL", err)
	}
	defer rows.Close()

	var urls []schemas.URL
	for rows.Next() {
		var url schemas.URL
		if err := rows.Scan(&url.Cod, &url.URL); err != nil {
			fmt.Println(err)
		}
		urls = append(urls, url)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

	return urls, nil

}

func GetUrl(cod string) (schemas.URL, error) {
	db, err := database.ConnectDatabse()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "select codigo, url_original  from encurtador_url where codigo = $1"

	var url schemas.URL

	err = db.QueryRow(query, cod).Scan(&url.Cod, &url.URL)
	if err != nil {
		log.Fatal("Error in get URL", err)
	}

	return url, nil
}

func InsertUrl(url string) (string, error) {
	db, err := database.ConnectDatabse()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	codUrl, err := usecase.Random()
	if err != nil {
		log.Fatal("Error generating URL", err.Error())
	}

	query := "insert into encurtador_url (codigo, url_original) values ($1, $2);"
	db.Exec(query, codUrl, url)

	result := fmt.Sprintf("Codigo encurtado: %s", codUrl)

	return result, nil
}
