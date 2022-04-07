package handlers

import (
	"belanjayukid_go/params"
	"belanjayukid_go/services"
	"net/http"
)

func HandleCreateProductUnit(w http.ResponseWriter, r *http.Request) {
	reqBody := params.ProductUnitRequest{}

	err := BindJSON(r, &reqBody)
	if err != nil {
		ToJSON(w, http.StatusBadRequest, badRequestResponse)
		return
	}

	response := services.CreateProductUnit(&reqBody)
	ToJSON(w, response.HttpCode, response)
}

func HandleGetProductUnitList(w http.ResponseWriter, r *http.Request) {

	response := services.GetProductUnitList()
	ToJSON(w, response.HttpCode, response)
}