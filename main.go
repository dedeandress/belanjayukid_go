package main

import (
	"go_sample_login_register/http/server"
	"go_sample_login_register/repositories"
	"go_sample_login_register/validators"

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
