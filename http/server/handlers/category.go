package handlers

import (
	"belanjayukid_go/params"
	"belanjayukid_go/services"
	"net/http"
)

func HandleCreateCategory(w http.ResponseWriter, r *http.Request) {
	reqBody := params.CategoryRequest{}

	err := BindJSON(r, &reqBody)
	if err != nil {
		ToJSON(w, http.StatusBadRequest, badRequestResponse)
		return
	}

	response := services.CreateCategory(&reqBody)
	ToJSON(w, response.HttpCode, response)
}

func HandleGetCategoryList(w http.ResponseWriter, r *http.Request) {

	response := services.GetCategoryList()
	ToJSON(w, response.HttpCode, response)
}