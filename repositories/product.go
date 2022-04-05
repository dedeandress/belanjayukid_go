package repositories

import (
	"belanjayukid_go/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductDetailByProductDetailID(productDetailID string)(productDetail *models.ProductDetail, err error)
}

type productRepository struct {
	db *DataSource
}

var productRepo *productRepository

func GetProductRepository() ProductRepository{
	if DBTrx != nil {
		productRepo = &productRepository{db: DBTrx}
	}else {
		productRepo = &productRepository{db: DB}
	}

	return productRepo
}


func (p productRepository) GetProductDetailByProductDetailID(productDetailID string) (productDetail *models.ProductDetail, err error) {
	productDetail = &models.ProductDetail{}
	res := productRepo.db.Preload("Product").Preload("ProductUnit").Scopes(filterByProductDetailID(productDetailID)).Find(&productDetail)
	if res.Error != nil{
		return nil, err
	}
	return productDetail, err
}

func filterByProductDetailID(productDetailID string) func(db *gorm.DB) *gorm.DB {
	return makeFilterFunc("product_details.id = ?", productDetailID)
}