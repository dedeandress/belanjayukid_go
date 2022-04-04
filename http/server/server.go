package server

import (
	"fmt"
	"net/http"
	"os"
)

func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	handler := buildRouteHandler()
	fmt.Println("Listening on localhost:3000")
	http.ListenAndServe(":"+port, handler)
}
