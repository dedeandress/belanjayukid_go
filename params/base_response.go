package params

import (
	"go_sample_login_register/enums"
	"net/http"
)

const (
	SUCCESS_STATUS = "success"
	ERROR_STATUS   = "error"
)

type Response struct {
	Status     string        `json:"status"`
	HttpCode   int           `json:"-"`
	Payload    interface{}   `json:"payload,omitempty"`
	Message    *errorContext `json:"message,omitempty"`
	Pagination *Pagination   `json:"pagination,omitempty"`
	Result     interface{}   `json:"result,omitempty"`
}

type errorContext struct {
	Status string `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type ActionForm struct {
	ActionType     string `json:"action_type"`
	RedirectionUrl string `json:"redirection_url"`
}

type Result struct {
	ResultCode    string `json:"resultCode"`
	ResultStatus  string `json:"resultStatus"`
	ResultMessage string `json:"resultMessage"`
}

type Pagination struct {
	PageNum   int  `json:"page_num"`
	PageSize  *int `json:"page_size,omitempty"`
	PageCount int  `json:"page_count"`
}

func NewSuccessResponse(payload interface{}) Response {
	return Response{
		Status:   SUCCESS_STATUS,
		HttpCode: http.StatusOK,
		Payload:  payload,
	}
}

func NewErrorResponse(resultCode enums.ResultCode) Response {
	return Response{
		Status:   ERROR_STATUS,
		HttpCode: resultCode.HttpStatusCode(),
		Message: &errorContext{
			Status: resultCode.String(),
			Title:  resultCode.HttpStatusText(),
			Detail: resultCode.Description(),
		},
	}
}

func NewDANAResponse(result interface{}) Response {
	return Response{
		Status:   SUCCESS_STATUS,
		HttpCode: http.StatusOK,
		Result:   result,
	}
}
