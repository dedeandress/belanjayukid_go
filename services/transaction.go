package services

import (
	"belanjayukid_go/enums"
	"belanjayukid_go/models"
	"belanjayukid_go/params"
	"belanjayukid_go/repositories"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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

	totalPrice := decimal.NewFromInt(0)

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

		if availableStock >= numberOfPurchases {
			subTotalPrice := productDetail.SellingPrice.Mul(decimal.NewFromInt(int64(trxDetails.NumberOfPurchases)))
			totalPrice = totalPrice.Add(subTotalPrice)
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

	//update transaction status to 1 (IN_PROGRESS)
	err = transactionRepo.UpdateTrxStatus(request.TransactionID, enums.IN_PROGRESS_TRANSACTION)
	if err != nil{
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	//update transaction total price
	err = transactionRepo.UpdateTrxTotalPrice(request.TransactionID, totalPrice)
	if err != nil{
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	response := params.TransactionResponse{
		TransactionID: request.TransactionID,
		TransactionDetails: &transactionDetailListResponse,
		TotalPrice: totalPrice,
	}

	return createResponseSuccess(ResponseService{
		Payload:  response,
		CommitDB: true,
	})
}

func FinishTransaction(transactionID string) params.Response{
	repositories.BeginTransaction()
	transactionRepo := repositories.GetTransactionRepository()
	productRepo := repositories.GetProductRepository()

	//final check stock
	finishTransactionResponse := params.FinishTransactionResponse{
		TransactionID: transactionID,
	}
	productOutOfStockList := make([]params.ProductOutOfStock, 0)
	transactionDetails, err := transactionRepo.GetTransactionDetailByTransactionID(transactionID)
	if err != nil {
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	for _, trxDetail := range *transactionDetails {
		productDetail, err := productRepo.GetProductDetailByProductDetailID(trxDetail.ProductDetailID.String())
		if err != nil {
			return createResponseError(
				ResponseService{
					RollbackDB: true,
					Error:      err,
					ResultCode: enums.INTERNAL_SERVER_ERROR,
				})
		}

		fmt.Printf("%s %s %d\n", productDetail.ID, productDetail.Product.Name, productDetail.Product.Stock)

		numberOfPurchases := trxDetail.NumberOfPurchases * productDetail.QuantityPerUnit
		availableStock := productDetail.Product.Stock / productDetail.QuantityPerUnit

		if availableStock < numberOfPurchases {
			productDetailOutOfStockList := make([]params.ProductDetailOutOfStock, 0)
			//TODO: search all available stock on product detail
			productDetailOutOfStock := params.ProductDetailOutOfStock{
				ProductUnit: productDetail.ProductUnit.Name,
				AvailableStock: availableStock,
			}

			productDetailOutOfStockList = append(productDetailOutOfStockList, productDetailOutOfStock)

			productID := productDetail.ProductID.String()
			productOutOfStockList = append(productOutOfStockList, params.ProductOutOfStock{
				Name: &productDetail.Product.Name,
				ProductID: &productID,
				AvailableStock: productDetail.Product.Stock,
				Detail: productDetailOutOfStockList,
			})
		}
	}

	if len(productOutOfStockList) != 0 {
		finishTransactionResponse.ProductOutOfStock = &productOutOfStockList
		return createResponseError(ResponseService{
			Payload: finishTransactionResponse,
			ResultCode: enums.PRODUCT_OUT_OF_STOCK,
			RollbackDB: true,
		})
	}

	//do decrease product stock based on product on transaction detail list
	for _, trxDetail := range *transactionDetails{
		productDetail, err := productRepo.GetProductDetailByProductDetailID(trxDetail.ProductDetailID.String())
		if err != nil {
			return createResponseError(
				ResponseService{
					RollbackDB: true,
					Error:      err,
					ResultCode: enums.INTERNAL_SERVER_ERROR,
				})
		}

		fmt.Printf("%s %s %d\n", productDetail.ID, productDetail.Product.Name, productDetail.Product.Stock)

		numberOfPurchases := trxDetail.NumberOfPurchases * productDetail.QuantityPerUnit
		currentStock := productDetail.Product.Stock
		updateStock := currentStock - numberOfPurchases

		err = productRepo.UpdateStock(productDetail.ProductID.String(), updateStock)
		if err != nil {
			return createResponseError(
				ResponseService{
					RollbackDB: true,
					Error:      err,
					ResultCode: enums.INTERNAL_SERVER_ERROR,
				})
		}
	}

	//update transaction status to 2 (FINISH_TRANSACTION)
	err = transactionRepo.UpdateTrxStatus(transactionID, enums.FINISH_TRANSACTION)
	if err != nil{
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	return createResponseSuccess(ResponseService{
		Payload: finishTransactionResponse,
		CommitDB: true,
	})
}

func GetTransactionList(request params.GetTransactionListRequest) params.Response{
	transactionRepo := repositories.GetTransactionRepository()

	transactions, err := transactionRepo.GetTransactionList(request.TransactionID, request.Status)
	if err != nil {
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	transactionsResponse := make([]params.TransactionResponse, 0)
	for _, transaction := range *transactions {
		trxDate := transaction.Date

		transactionResponse := params.TransactionResponse{
			TransactionID: transaction.ID.String(),
			TotalPrice: transaction.TotalPrice,
			Date: &trxDate,
		}

		transactionsResponse = append(transactionsResponse, transactionResponse)
	}

	transactionListResponse := params.TransactionListResponse{
		Transactions: transactionsResponse,
	}

	return createResponseSuccess(ResponseService{
		Payload: transactionListResponse,
	})
}

func GetTransactionDetail(transactionID string) params.Response {
	transactionRepo := repositories.GetTransactionRepository()
	productRepo := repositories.GetProductRepository()

	transaction, err := transactionRepo.GetTransaction(transactionID)
	if err != nil {
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	transactionDetails, err := transactionRepo.GetTransactionDetailByTransactionID(transactionID)
	if err != nil {
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error:      err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	transactionDetailsResponse := make([]params.TransactionDetailResponse, 0)
	for _, trxDetails := range *transactionDetails {
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
			Product: params.ProductDetailResponse{
				Name: productDetail.Product.Name,
				SKU: productDetail.Product.SKU,
			},
			ProductUnit: productDetail.ProductUnit.Name,
			NumberOfPurchases: trxDetails.NumberOfPurchases,
		}

		transactionDetailsResponse = append(transactionDetailsResponse, transactionDetailResponse)
	}

	transactionResponse := params.TransactionResponse{
		TransactionID: transactionID,
		Date: &transaction.Date,
		TotalPrice: transaction.TotalPrice,
		TransactionDetails: &transactionDetailsResponse,
	}

	return createResponseSuccess(ResponseService{
		Payload: transactionResponse,
	})
}

func mapToModels(request *params.TransactionRequest) ([]models.TransactionDetail, error){
	transactionID, err := uuid.Parse(request.TransactionID)
	if err != nil {
		return nil, err
	}

	transactionDetails := make([]models.TransactionDetail, 0)
	for _, details := range request.TransactionDetails{
		transactionDetailID := uuid.New()
		if details.TransactionDetailID != nil {
			transactionDetailID, err = uuid.Parse(*details.TransactionDetailID)
			if err != nil {
				return nil, err
			}
		}
		productDetailID, err := uuid.Parse(details.ProductDetailID)
		if err != nil {
			return nil, err
		}
		transactionDetail := models.TransactionDetail{ID: transactionDetailID, TransactionID: transactionID, NumberOfPurchases: details.NumberOfPurchases, ProductDetailID: productDetailID}
		transactionDetails = append(transactionDetails, transactionDetail)
	}

	return transactionDetails, nil
}