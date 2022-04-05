package handlers

import (
	"belanjayukid_go/crypto"
	"belanjayukid_go/params"
	"context"
	"net/http"
)

func Auth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	preprocessedHandler := func(responseWriter http.ResponseWriter, request *http.Request) {
		payload, resultCode, err := crypto.GetHttpRequestAuthorizationClaim(request)
		if err != nil {
			ToJSON(responseWriter, resultCode.HttpStatusCode(), params.NewErrorResponse(resultCode, nil))
			return
		}

		contextUser := context.WithValue(request.Context(), CONTEXT_USER, payload)
		request = request.WithContext(contextUser)
		handlerFunc(responseWriter, request)
	}
	return preprocessedHandler
}
