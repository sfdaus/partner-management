package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// CreatePartnerReq represent create request body
type CreatePartnerReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (request CreatePartnerReq) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Name, validation.Required),
	)
}

// Update request body
type UpdatePartnerReq struct {
	ID          string `param:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    *bool  `json:"is_active"`
}

func (request UpdatePartnerReq) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.ID, validation.Required),
	)
}

// Delete request body
type DeletePartnerReq struct {
	ID string `param:"id"`
}

func (request DeletePartnerReq) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.ID, validation.Required),
	)
}

// GetList request body
type GetListPartnerReq struct {
	Name     string `query:"name"`
	IsActive *bool  `query:"is_active"`
	PerPage  int64  `query:"per_page"`
	Page     int64  `query:"page"`
}

func (request GetListPartnerReq) Validate() error {
	return validation.ValidateStruct(
		&request,
	)
}

// GetDetail request body
type GetDetailPartnerReq struct {
	ID string `param:"id"`
}

func (request GetDetailPartnerReq) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.ID, validation.Required),
	)
}
