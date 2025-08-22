package request

import validation "github.com/go-ozzo/ozzo-validation"

type GetListCompensationTypesReq struct {
	Name     string `query:"name"`
	IsActive *bool  `query:"is_active"`
	PerPage  int64  `query:"per_page"`
	Page     int64  `query:"page"`
}

func (request GetListCompensationTypesReq) Validate() error {
	return validation.ValidateStruct(
		&request,
	)
}

type CreateCompensationTypeReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string
}

func (request CreateCompensationTypeReq) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.UserID, validation.Required),
	)
}
