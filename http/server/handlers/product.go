package handlers

import (
	"belanjayukid_go/services"
	"net/http"
)

func HandleGetProductList(w http.ResponseWriter, r *http.Request) {

	response := services.GetProductList()
	ToJSON(w, response.HttpCode, response)
}