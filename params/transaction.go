package params

import (
	"belanjayukid_go/validators"
	"github.com/shopspring/decimal"
	"time"
)

type TransactionRequest struct {
	TransactionID string `json:"transaction_id"`
	TransactionDetails []transactionDetailRequest `json:"transaction_details"`
}

type transactionDetailRequest struct {
	TransactionDetailID *string `json:"transaction_detail_id"`
	ProductDetailID string `json:"product_detail_id"`
	NumberOfPurchases int `json:"number_of_purchases"`
}

type GetTransactionListRequest struct {
	TransactionID *string `json:"transaction_id"`
	Status *int `json:"status"`
}

func (request TransactionRequest) Validate() error {
	return validators.ValidateInputs(request)
}

func (request GetTransactionListRequest) Validate() error {
	return validators.ValidateInputs(request)
}

type TransactionResponse struct {
	TransactionID string `json:"transaction_id"`
	TransactionDetails *[]TransactionDetailResponse `json:"transaction_details"`
	TotalPrice decimal.Decimal `json:"total_price"`
	Date *time.Time `json:"date"`
}

type TransactionListResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
}

type TransactionDetailResponse struct {
	NumberOfPurchases int `json:"number_of_purchases"`
	Product TransactionProductDetailResponse `json:"product"`
	ProductUnit string `json:"product_unit"`
	ProductOutOfStock *ProductOutOfStock `json:"product_out_of_stock"`
}

type FinishTransactionResponse struct {
	TransactionID string `json:"transaction_id"`
	ProductOutOfStock *[]ProductOutOfStock `json:"product_out_of_stock"`
}

type TransactionProductDetailResponse struct {
	SKU string
	Name string
}

type ProductOutOfStock struct {
	ProductID *string `json:"product_id"`
	Name *string `json:"name"`
	AvailableStock int `json:"available_stock"`
	Detail []ProductDetailOutOfStock `json:"detail"`
}

type ProductDetailOutOfStock struct {
	ProductUnit string `json:"product_unit"`
	AvailableStock int `json:"available_stock"`
}