package models

import (
	"encurtador-url/cmd/database"
	"encurtador-url/internal/usecase"
	"fmt"
	"log"
)

type Encurtador struct {
	ID          int    `json:"id"`
	Codigo      string `json:"codigo"`
	UrlOriginal string `json:"urloriginal"`
}

type Asdf struct {
	Codigo string `json:"codigo"`
}

func InsertUrl(url string) (string, error) {
	db, err := database.ConnectionDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	codigoUrl, err := usecase.RandomUrl(6)
	if err != nil {
		log.Fatal("Falha ao gerar o codigo Random para URL", err.Error())
	}

	insertSQL := `
		INSERT INTO encurtador_url (codigo, url_original)
		VALUES($1, $2)
	`
	db.Exec(insertSQL, codigoUrl, url)

	fmt.Println(codigoUrl)

	result := fmt.Sprintf("Resultado da URL = %s", codigoUrl)

	return result, nil

}

type UrlOriginal struct {
	UrlOrignal string `json:"urloriginal"`
}

func GetIDUrl(codigo string) ([]UrlOriginal, error) {
	db, err := database.ConnectionDB()
	if err != nil {
		fmt.Println("Falha ao conectar com o banco de dados", err.Error())
	}

	defer db.Close()

	query := `
		SELECT 
			url_original
		FROM encurtador_url 
		WHERE codigo = $1
	`
	rows, err := db.Query(query, codigo)
	if err != nil {
		fmt.Println("Erro ao consultar no banco de dados", err.Error())
	}
	defer rows.Close()

	var urlOriginal []UrlOriginal
	for rows.Next() {
		var url string

		err := rows.Scan(&url)
		if err != nil {
			fmt.Println("Erro ao buscar no banco de dados", err.Error())
		}

		urlOriginal = append(urlOriginal, UrlOriginal{UrlOrignal: url})
	}

	return urlOriginal, nil
}
