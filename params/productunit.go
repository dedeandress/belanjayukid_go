package params

import "belanjayukid_go/validators"

type ProductUnitRequest struct {
	Name string `json:"name"`
}

func (request ProductUnitRequest) Validate() error {
	return validators.ValidateInputs(request)
}

type ProductUnitListResponse struct {
	ProductUnits []ProductUnitResponse
}

type ProductUnitResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
}
