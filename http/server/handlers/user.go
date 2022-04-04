package handlers

import (
	"belanjayukid_go/crypto"
	"belanjayukid_go/params"
	"belanjayukid_go/services"
	"net/http"
)

func HandleLogin(responseWriter http.ResponseWriter, request *http.Request) {
	reqBody := params.LoginRequest{}

	err := BindJSON(request, &reqBody)
	if err != nil {
		ToJSON(responseWriter, http.StatusBadRequest, badRequestResponse)
		return
	}

	response := services.Login(&reqBody)
	ToJSON(responseWriter, response.HttpCode, response)
}

func HandleGetMe(responseWriter http.ResponseWriter, request *http.Request) {
	payload := request.Context().Value(CONTEXT_USER).(*crypto.Payload)
	response := services.GetMe(payload.UserID)
	ToJSON(responseWriter, response.HttpCode, response)
}

func HandleRegister(responseWriter http.ResponseWriter, request *http.Request) {
	reqBody := params.RegisterRequest{}

	err := BindJSON(request, &reqBody)
	if err != nil {
		ToJSON(responseWriter, http.StatusBadRequest, badRequestResponse)
		return
	}

	response := services.Register(&reqBody)
	ToJSON(responseWriter, response.HttpCode, response)
}
