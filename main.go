package main

import (
	"belanjayukid_go/http/server"
	"belanjayukid_go/repositories"
	"belanjayukid_go/validators"

	"log"
)

func main() {
	err := repositories.InitDBFactory()
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = validators.InitValidator()
	if err != nil {
		log.Fatalln(err)
		return
	}

	server.StartServer()
}
