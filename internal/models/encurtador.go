package models

import (
	"encurtador-url/cmd/config"
	"encurtador-url/cmd/database"
	"encurtador-url/internal/usecase"
	"fmt"
	"log"
	"strconv"
)

type Encurtador struct {
	ID          int    `json:"id"`
	Codigo      string `json:"codigo"`
	UrlOriginal string `json:"urloriginal"`
}

func InsertUrl(url string) (string, error) {
	db, err := database.ConnectionDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	lengthRandomUrl, err := strconv.Atoi(config.LoadEnv().LengthUrl)
	if err != nil {
		fmt.Println("Erro ao converter o valor", err.Error())
	}

	codigoUrl, err := usecase.RandomUrl(lengthRandomUrl)
	if err != nil {
		log.Fatal("Falha ao gerar o codigo Random para URL", err.Error())
	}

	insertSQL := `
		INSERT INTO encurtador_url (codigo, url_original)
		VALUES($1, $2)
	`

	var insertCode string
	err2 := db.QueryRow(insertSQL, codigoUrl, url).Scan(&insertCode)
	if err != nil {
		return "", err2
	}

	fmt.Printf("Codigo gerado e inserido no banco de dados: %s\n", insertCode)
	return insertCode, nil

}

func GetAllUrl() ([]Encurtador, error) {
	db, err := database.ConnectionDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	consulta := "SELECT id, codigo, url_original FROM encurtador_url ORDER BY ID"

	rows, err := db.Query(consulta)
	if err != nil {
		fmt.Println("Erro ao consultar no banco de dados", err.Error())
	}

	defer rows.Close()

	var e []Encurtador
	for rows.Next() {
		var id int
		var codigo string
		var urlOrignal string

		err := rows.Scan(&id, &codigo, &urlOrignal)
		if err != nil {
			fmt.Println("Erro ao buscar no banco de dados", err.Error())
		}

		e = append(e, Encurtador{ID: id, Codigo: codigo, UrlOriginal: urlOrignal})
	}

	return e, nil
}
