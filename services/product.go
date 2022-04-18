package services

import (
	"belanjayukid_go/enums"
	"belanjayukid_go/params"
	"belanjayukid_go/repositories"
)

func GetProductList() params.Response {
	productRepo := repositories.GetProductRepository()

	products, err := productRepo.GetProductList()
	if err != nil {
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error: err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	productListResponse := make([]params.ProductResponse, 0)
	for _, product := range products {
		categoryResponse := params.CategoryResponse{
			ID: product.Category.ID.String(),
			Name: product.Category.Name,
		}
		productResponse := params.ProductResponse{
			ID: product.ID.String(),
			Name: product.Name,
			SKU: product.SKU,
			Category: categoryResponse,
			Stock: product.Stock,
			ImageURL: product.ImageURL,
		}

		productDetailListResponse := make([]params.ProductDetailResponse, 0)
		for _, productDetail := range product.ProductDetails{
			productDetailResponse := params.ProductDetailResponse{
				ID: productDetail.ID.String(),
				SellingPrice: productDetail.SellingPrice,
				PurchasesPrice: productDetail.PurchasePrice,
				QuantityPerUnit: productDetail.QuantityPerUnit,
				Unit: productDetail.ProductUnit.Name,
			}

			productDetailListResponse = append(productDetailListResponse, productDetailResponse)
		}

		productResponse.ProductDetail = &productDetailListResponse

		productListResponse = append(productListResponse, productResponse)
	}

	return createResponseSuccess(ResponseService{Payload: productListResponse})
}
