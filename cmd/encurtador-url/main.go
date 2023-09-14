package main

import (
	"encurtador-url/internal/usecase"
	"encurtador-url/pkg/server"
)

func main() {

	usecase.CreateTable()

	s := server.NewServer()
	s.Run()

}
