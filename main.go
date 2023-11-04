package main

import (
	"encurtador-url/routes"
	"encurtador-url/usecase"
)

func main() {
	usecase.CreateTable()
	routes.Inicialize()
}
