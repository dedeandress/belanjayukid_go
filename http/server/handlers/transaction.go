package handlers

import (
	"belanjayukid_go/params"
	"belanjayukid_go/services"
	"github.com/gorilla/mux"
	"net/http"
)

func HandleAddToCartTransaction(w http.ResponseWriter, r *http.Request) {
	reqBody := params.TransactionRequest{}

	err := BindJSON(r, &reqBody)
	if err != nil {
		ToJSON(w, http.StatusBadRequest, badRequestResponse)
		return
	}

	response := services.AddToCart(&reqBody)
	ToJSON(w, response.HttpCode, response)
}

func HandleInitTransaction(w http.ResponseWriter, r *http.Request) {
	response := services.InitTransaction()
	ToJSON(w, response.HttpCode, response)
}

func HandleFinishTransaction(w http.ResponseWriter, r *http.Request) {
	muxParams := mux.Vars(r)
	transactionID := muxParams["transactionID"]

	response := services.FinishTransaction(transactionID)
	ToJSON(w, response.HttpCode, response)
}