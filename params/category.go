package params

import "belanjayukid_go/validators"

type CategoryRequest struct {
	Name string `json:"name"`
}

func (request CategoryRequest) Validate() error {
	return validators.ValidateInputs(request)
}

type CategoryListResponse struct {
	Categories []CategoryResponse
}

type CategoryResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
}
