package params

import "github.com/shopspring/decimal"

type ProductRequest struct {
	ProductID string `json:"product_id"`
}

type ProductResponse struct {
	ID string `json:"id"`
	SKU string `json:"sku"`
	Name string `json:"name"`
	Stock int `json:"stock"`
	Category CategoryResponse `json:"category"`
	ImageURL string `json:"image_url"`
	ProductDetail *[]ProductDetailResponse `json:"product_details"`
}

type ProductDetailResponse struct {
	ID string `json:"id"`
	Unit string `json:"unit"`
	SellingPrice decimal.Decimal `json:"selling_price"`
	PurchasesPrice decimal.Decimal `json:"purchases_price"`
	QuantityPerUnit int `json:"quantity_per_unit"`
}