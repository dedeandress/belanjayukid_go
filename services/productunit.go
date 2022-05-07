package services

import (
	"belanjayukid_go/enums"
	"belanjayukid_go/models"
	"belanjayukid_go/params"
	"belanjayukid_go/repositories"
)

func CreateProductUnit(request *params.ProductUnitRequest) params.Response {
	repositories.BeginTransaction()
	productUnitRepo := repositories.GetProductUnitRepository()

	productUnit, err := productUnitRepo.Insert(&models.ProductUnit{Name: request.Name})
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
			ID: productUnit.ID.String(),
			Name: productUnit.Name,
		},
	})
}

func GetProductUnitList() params.Response {
	productUnitRepo := repositories.GetProductUnitRepository()

	productUnitList, err := productUnitRepo.GetProductUnitList()
	if err != nil {
		return createResponseError(ResponseService{
			Error: err,
			ResultCode: enums.INTERNAL_SERVER_ERROR,
		})
	}

	productUnitResponseList := make([]params.ProductUnitResponse, 0)
	for _, category := range *productUnitList {
		productUnitResponse := params.ProductUnitResponse{
			ID: category.ID.String(),
			Name: category.Name,
		}

		productUnitResponseList = append(productUnitResponseList, productUnitResponse)
	}

	return createResponseSuccess(ResponseService{
		Payload: params.ProductUnitListResponse{ProductUnits: productUnitResponseList},
	})
}
