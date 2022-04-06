package services

import (
	"belanjayukid_go/enums"
	"belanjayukid_go/models"
	"belanjayukid_go/params"
	"belanjayukid_go/repositories"
	"github.com/google/uuid"
)

func CreateCategory(request *params.CategoryRequest) params.Response {
	repositories.BeginTransaction()
	categoryRepo := repositories.GetCategoryRepository()

	category, err := categoryRepo.Insert(models.Category{ID: uuid.New(), Name: request.Name})
	if err != nil {
		return createResponseError(
			ResponseService{
				RollbackDB: true,
				Error: err,
				ResultCode: enums.INTERNAL_SERVER_ERROR,
			})
	}

	return createResponseSuccess(ResponseService{
		CommitDB: true,
		Payload: params.CategoryResponse{
			ID: category.ID.String(),
			Name: category.Name,
		},
	})
}

func GetCategoryList() params.Response {
	categoryRepo := repositories.GetCategoryRepository()

	categoryList, err := categoryRepo.GetCategoryList()
	if err != nil {
		return createResponseError(ResponseService{
			Error: err,
			ResultCode: enums.INTERNAL_SERVER_ERROR,
		})
	}

	categoryResponseList := make([]params.CategoryResponse, 0)
	for _, category := range *categoryList {
		categoryResponse := params.CategoryResponse{
			ID: category.ID.String(),
			Name: category.Name,
		}

		categoryResponseList = append(categoryResponseList, categoryResponse)
	}

	return createResponseSuccess(ResponseService{
		Payload: params.CategoryListResponse{Categories: categoryResponseList},
	})
}
