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
