package services

import (
	"belanjayukid_go/enums"
	"belanjayukid_go/models"
	"belanjayukid_go/params"
	"belanjayukid_go/repositories"
	"fmt"
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
	productRepo := repositories.GetProductRepository()

	transactionDetailInsertDBList, err := mapToModels(request)
	if err != nil {
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	// check stock is exist and mapping to transactionDetailResponse
	transactionDetailListResponse := make([]params.TransactionDetailResponse, 0)
	for index, trxDetails := range transactionDetailInsertDBList {

		productDetail, err := productRepo.GetProductDetailByProductDetailID(trxDetails.ProductDetailID.String())
		if err != nil {
			return createResponseError(
				ResponseService{
					RollbackDB: true,
					Error:      err,
					ResultCode: enums.INTERNAL_SERVER_ERROR,
				})
		}
		transactionDetailResponse := params.TransactionDetailResponse{
			NumberOfPurchases: trxDetails.NumberOfPurchases,
			Product: params.ProductDetailResponse{
				Name: productDetail.Product.Name,
				SKU: productDetail.Product.SKU,
			},
			ProductUnit: productDetail.ProductUnit.Name,
		}

		fmt.Printf("%s %s %d\n", productDetail.ID, productDetail.Product.Name, productDetail.Product.Stock)

		numberOfPurchases := trxDetails.NumberOfPurchases * productDetail.QuantityPerUnit
		availableStock := productDetail.Product.Stock / productDetail.QuantityPerUnit

		if availableStock < numberOfPurchases {
			productDetailOutOfStockList := make([]params.ProductDetailOutOfStock, 0)
			//TODO: search all available stock on product detail
			productDetailOutOfStock := params.ProductDetailOutOfStock{
				ProductUnit: productDetail.ProductUnit.Name,
				AvailableStock: availableStock,
			}

			productDetailOutOfStockList = append(productDetailOutOfStockList, productDetailOutOfStock)

			transactionDetailResponse.ProductOutOfStock = &params.ProductOutOfStock{
				AvailableStock: productDetail.Product.Stock,
				Detail: productDetailOutOfStockList,
			}

			//delete the out-of-stock product from transaction details insert db list
			transactionDetailInsertDBList = append(transactionDetailInsertDBList[:index], transactionDetailInsertDBList[index+1:]...)
		}

		transactionDetailListResponse = append(transactionDetailListResponse, transactionDetailResponse)
	}

	// insert the transaction detail
	err = transactionRepo.Update(transactionDetailInsertDBList)
	if err != nil {
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	response := params.TransactionResponse{
		TransactionID: request.TransactionID,
		TransactionDetails: transactionDetailListResponse,
	}

	return createResponseSuccess(ResponseService{
		Payload:  response,
		CommitDB: true,
	})
}

func finishTransaction() {

	//final check stock

	//if still have product out of stock throw error list of product that are out of stock

	//if don't have product out of stock
	//do decrease product stock based on product on transaction detail list

	//update transaction status
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