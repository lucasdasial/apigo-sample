package main

import "github.com/luccasalves/apigo-sample/server"

func main() {
	server := server.CreateServer()

	server.Run()
}
