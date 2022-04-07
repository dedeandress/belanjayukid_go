package repositories

import "belanjayukid_go/models"

type ProductUnitRepository interface {
	Insert(productUnit models.ProductUnit)(insertedProductUnit *models.ProductUnit, err error)
	GetProductUnitList()(*[]models.ProductUnit, error)
}

type productUnitRepository struct {
	db *DataSource
}

var productUnitRepo *productUnitRepository

func GetProductUnitRepository() ProductUnitRepository {
	if DBTrx != nil {
		productUnitRepo = &productUnitRepository{db: DBTrx}
	} else {
		productUnitRepo = &productUnitRepository{db: DB}
	}

	return productUnitRepo
}

func (c productUnitRepository) Insert(productUnit models.ProductUnit) (insertedProductUnit *models.ProductUnit, err error) {
	insertedProductUnit = &models.ProductUnit{}
	res := productUnitRepo.db.Create(productUnit).Scan(insertedProductUnit)
	if res.Error != nil {
		return nil, res.Error
	}

	return insertedProductUnit, nil
}

func (c productUnitRepository) GetProductUnitList() (*[]models.ProductUnit, error) {
	var productUnitList *[]models.ProductUnit
	res := productUnitRepo.db.Find(&productUnitList)
	if res.Error != nil {
		return nil, res.Error
	}

	return productUnitList, nil
}
