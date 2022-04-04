package handlers

import (
	"encoding/json"
	"go_sample_login_register/enums"
	"go_sample_login_register/params"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var badRequestResponse params.Response
var internalServerErrorResponse params.Response
var decoder *schema.Decoder

const (
	HEADER_JSON_TYPE = "application/json"
	CONTEXT_USER     = "context_user"
)

func init() {
	badRequestResponse = params.NewErrorResponse(enums.BAD_REQUEST)
	internalServerErrorResponse = params.NewErrorResponse(enums.INTERNAL_SERVER_ERROR)
	decoder = schema.NewDecoder()
}

func ToJSON(w http.ResponseWriter, statusCode int, obj interface{}) {
	w.Header().Add("Content-Type", HEADER_JSON_TYPE)
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(&obj)
}

func BindJSON(req *http.Request, request params.RequestParams) error {
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(request); err != nil {
		return err
	}
	return request.Validate()
}

func readQueryURLVariables(request *http.Request, requestBody params.RequestParams) error {
	vars := make(map[string][]string)
	for key, value := range mux.Vars(request) {
		vars[key] = []string{value}
	}
	err := decoder.Decode(requestBody, vars)
	if err != nil {
		return err
	}
	return requestBody.Validate()
}

func decodeQueryURL(dest interface{}, request *http.Request) error {
	return decoder.Decode(dest, request.URL.Query())
}
