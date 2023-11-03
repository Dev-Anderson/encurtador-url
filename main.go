package main

import (
	"encurtador-url/database"
	"encurtador-url/routes"
)

func main() {
	database.ConnectDatabse()
	routes.Inicialize()
}
