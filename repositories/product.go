package repositories

import (
	"belanjayukid_go/models"
	"github.com/kr/pretty"
	"gorm.io/gorm"
	"log"
)

type ProductRepository interface {
	GetProductList()([]models.Product, error)
	GetProductDetailByProductDetailID(productDetailID string)(productDetail *models.ProductDetail, err error)
	UpdateStock(productID string, stock int)(err error)
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


func (p productRepository) GetProductList() ([]models.Product, error) {
	var productList []models.Product
	res := productRepo.db.Preload("ProductDetails").Preload("ProductDetails.ProductUnit").Preload("Category").Find(&productList)
	if res.Error != nil {
		return nil, res.Error
	}

	log.Print(pretty.Sprint(productList))

	return productList, nil
}

func (p productRepository) GetProductDetailByProductDetailID(productDetailID string) (productDetail *models.ProductDetail, err error) {
	productDetail = &models.ProductDetail{}
	res := productRepo.db.Preload("Product").Preload("ProductUnit").Scopes(filterByProductDetailID(productDetailID)).Find(&productDetail)
	if res.Error != nil{
		return nil, err
	}
	return productDetail, err
}

func (p productRepository) UpdateStock(productID string, stock int) (err error) {
	err = productRepo.db.Model(models.Product{}).Where("id = ?", productID).Update("stock", stock).Error
	return err
}

func filterByProductDetailID(productDetailID string) func(db *gorm.DB) *gorm.DB {
	return makeFilterFunc("product_details.id = ?", productDetailID)
}