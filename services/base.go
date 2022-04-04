package services

import (
	"go_sample_login_register/enums"
	"go_sample_login_register/models"
	"go_sample_login_register/params"
	"go_sample_login_register/repositories"
)

type ResponseService struct {
	Payload    interface{}
	CommitDB   bool
	RollbackDB bool
	IsNotFound bool
	Error      error
	ResultCode enums.ResultCode
	Result     interface{}
}

func createResponseSuccess(response ResponseService) params.Response {
	if response.CommitDB {
		repositories.CommitTransaction()
	}
	return params.NewSuccessResponse(response.Payload)
}

func createResponseError(response ResponseService) params.Response {
	if response.RollbackDB {
		repositories.RollbackTransaction()
	}

	if response.IsNotFound {
		switch response.Payload.(type) {
		case *models.User:
			response.ResultCode = enums.USER_NOT_FOUND
		default:
			response.ResultCode = enums.INTERNAL_SERVER_ERROR
		}
	}

	return params.NewErrorResponse(response.ResultCode)
}
