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
