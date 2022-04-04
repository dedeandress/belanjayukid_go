package services

import (
	"go_sample_login_register/crypto"
	"go_sample_login_register/enums"
	"go_sample_login_register/models"
	"go_sample_login_register/params"
	"go_sample_login_register/repositories"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Register(request *params.RegisterRequest) params.Response {
	repositories.BeginTransaction()
	userRepo := repositories.GetUserRepository()

	_, err := userRepo.GetUserByEmail(request.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
			if err != nil {
				return createResponseError(
					ResponseService{
						RollbackDB: true,
						Error:      err,
						ResultCode: enums.INTERNAL_SERVER_ERROR,
					})
			}
			return insertNewUser(request.Email, string(hash))
		}
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	return createResponseError(ResponseService{
		Error:      err,
		ResultCode: enums.USER_ALREADY_EXIST,
	})
}

func Login(request *params.LoginRequest) params.Response {
	userRepo := repositories.GetUserRepository()
	user, err := userRepo.GetUserByEmail(request.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return createResponseError(
				ResponseService{
					Error:      err,
					ResultCode: enums.WRONG_EMAIL_PASSWORD,
				})
		}
		return createResponseError(
			ResponseService{
				Error:      err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return createResponseError(
			ResponseService{
				Error:      err,
				ResultCode: enums.WRONG_EMAIL_PASSWORD,
			})
	}

	token, err := crypto.GenerateJWT(user.ID.String())

	if err != nil {
		return serverErrorResponse(err)
	}

	return createResponseSuccess(ResponseService{
		Payload: params.LoginResponse{AuthToken: token},
	})
}

func GetMe(userID string) params.Response {
	userRepo := repositories.GetUserRepository()

	user, err := userRepo.GetUserByID(userID)
	if err != nil {
		return createResponseError(
			ResponseService{
				Error:      err,
				ResultCode: enums.BAD_REQUEST,
			})
	}

	return createResponseSuccess(
		ResponseService{
			Payload: params.MeResponse{
				Email: user.Email,
			},
		},
	)
}

func insertNewUser(email string, password string) params.Response {
	userRepo := repositories.GetUserRepository()

	insertedUser, err := userRepo.Insert(&models.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return serverErrorResponse(err)
	}

	token, err := crypto.GenerateJWT(insertedUser.ID.String())

	if err != nil {
		return serverErrorResponse(err)
	}

	return createResponseSuccess(ResponseService{
		Payload:  params.RegisterResponse{AuthToken: token, Email: insertedUser.Email},
		CommitDB: true,
	})
}

func serverErrorResponse(err error) params.Response {
	return createResponseError(
		ResponseService{
			RollbackDB: true,
			Error:      err,
			ResultCode: enums.INTERNAL_SERVER_ERROR,
		})
}
