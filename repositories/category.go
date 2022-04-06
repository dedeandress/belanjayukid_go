package repositories

import "belanjayukid_go/models"

type CategoryRepository interface {
	Insert(category models.Category)(insertedCategory *models.Category, err error)
	GetCategoryList()(*[]models.Category, error)
}

type categoryRepository struct {
	db *DataSource
}

var categoryRepo *categoryRepository

func GetCategoryRepository() CategoryRepository {
	if DBTrx != nil {
		categoryRepo = &categoryRepository{db: DBTrx}
	} else {
		categoryRepo = &categoryRepository{db: DB}
	}

	return categoryRepo
}

func (c categoryRepository) Insert(category models.Category) (insertedCategory *models.Category, err error) {
	insertedCategory = &models.Category{}
	res := categoryRepo.db.Create(category).Scan(insertedCategory)
	if res.Error != nil {
		return nil, res.Error
	}

	return insertedCategory, nil
}

func (c categoryRepository) GetCategoryList() (*[]models.Category, error) {
	var categoryList *[]models.Category
	res := categoryRepo.db.Find(&categoryList)
	if res.Error != nil {
		return nil, res.Error
	}

	return categoryList, nil
}
