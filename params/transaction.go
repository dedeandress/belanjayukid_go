package params

import "belanjayukid_go/validators"

type TransactionRequest struct {
	TransactionID string `json:"transaction_id"`
	TransactionDetails []transactionDetailRequest `json:"transaction_details"`
}

type transactionDetailRequest struct {
	ProductDetailID string `json:"product_detail_id"`
	NumberOfPurchases int `json:"number_of_purchases"`
}

func (request TransactionRequest) Validate() error {
	return validators.ValidateInputs(request)
}

type TransactionResponse struct {
	TransactionID string `json:"transaction_id"`
	TransactionDetails []TransactionDetailResponse `json:"transaction_details"`
}

type TransactionDetailResponse struct {
	NumberOfPurchases int
	Product ProductDetailResponse
	ProductUnit string
	ProductOutOfStock *ProductOutOfStock
}

type ProductDetailResponse struct {
	SKU string
	Name string
}

type ProductOutOfStock struct {
	AvailableStock int
	Detail []ProductDetailOutOfStock
}

type ProductDetailOutOfStock struct {
	ProductUnit string
	AvailableStock int
}