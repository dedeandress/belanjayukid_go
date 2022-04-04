package enums

import (
	"net/http"
	"strconv"
)

type ResultCode int

const (
	SUCCESS               ResultCode = 0
	NO_CONTENT            ResultCode = 1
	INTERNAL_SERVER_ERROR ResultCode = 1000
	INVALID_AUTH_TOKEN    ResultCode = 1001
	BAD_REQUEST           ResultCode = 1002
	PERMISSION_DENIED     ResultCode = 1003
	NOT_FOUND             ResultCode = 1004
	USER_NOT_FOUND        ResultCode = 1005
	USER_ALREADY_EXIST    ResultCode = 1006
	WRONG_EMAIL_PASSWORD  ResultCode = 1007
)

func (resultCode ResultCode) HttpStatusCode() int {
	switch resultCode {
	case SUCCESS:
		return http.StatusOK
	case NO_CONTENT:
		return http.StatusNoContent
	case BAD_REQUEST:
		return http.StatusBadRequest
	case INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case PERMISSION_DENIED:
		return http.StatusForbidden
	case INVALID_AUTH_TOKEN:
		return http.StatusUnauthorized
	case WRONG_EMAIL_PASSWORD:
		return http.StatusForbidden
	case USER_ALREADY_EXIST:
		return http.StatusConflict
	case USER_NOT_FOUND:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func (resultCode ResultCode) HttpStatusText() string {
	statusCode := resultCode.HttpStatusCode()
	return http.StatusText(statusCode)
}

func (resultCode ResultCode) String() string {
	return strconv.Itoa(int(resultCode))
}

func (resultCode ResultCode) Description() string {
	switch resultCode {
	case SUCCESS:
		return "Success"
	case BAD_REQUEST:
		return "Bad Request"
	case INTERNAL_SERVER_ERROR:
		return "Internal Server Error"
	case PERMISSION_DENIED:
		return "Permission denied"
	case INVALID_AUTH_TOKEN:
		return "Auth token is invalid or has expired"
	case USER_ALREADY_EXIST:
		return "User Already Exist"
	case USER_NOT_FOUND:
		return "User not found"
	case WRONG_EMAIL_PASSWORD:
		return "Wrong Email/Password"
	default:
		return "Internal Server Error"
	}
}
