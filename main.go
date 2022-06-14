package main

import (
	"github.com/luccasalves/apigo-sample/database"
	"github.com/luccasalves/apigo-sample/server"
)

func main() {
	database.StartDB()

	server := server.CreateServer()

	server.Run()
}
