package services

import (
	"belanjayukid_go/enums"
	"belanjayukid_go/models"
	"belanjayukid_go/params"
	"belanjayukid_go/repositories"
	"github.com/google/uuid"
)

func InitTransaction() params.Response {
	repositories.BeginTransaction()
	transactionRepo := repositories.GetTransactionRepository()
	transaction, err := transactionRepo.Insert()
	if err != nil {
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error: err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	return createResponseSuccess(ResponseService{
		Payload: transaction,
		CommitDB: true,
	})
}

func AddToCart(request *params.TransactionRequest) params.Response {
	repositories.BeginTransaction()
	transactionRepo := repositories.GetTransactionRepository()

	transactionDetails, err := mapToModels(request)
	if err != nil {
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	err = transactionRepo.Update(transactionDetails)
	if err != nil {
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	return createResponseSuccess(ResponseService{
		Payload:  transactionDetails,
		CommitDB: true,
	})
}

func mapToModels(request *params.TransactionRequest) ([]models.TransactionDetail, error){
	transactionID, err := uuid.Parse(request.TransactionID)
	if err != nil {
		return nil, err
	}

	transactionDetails := make([]models.TransactionDetail, 0)
	for _, details := range request.TransactionDetails{
		productDetailID, err := uuid.Parse(details.ProductDetailID)
		if err != nil {
			return nil, err
		}
		transactionDetail := models.TransactionDetail{ID: uuid.New(), TransactionID: transactionID, NumberOfPurchases: details.NumberOfPurchases, ProductDetailID: productDetailID}
		transactionDetails = append(transactionDetails, transactionDetail)
	}

	return transactionDetails, nil
}