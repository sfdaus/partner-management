package domain

import (
	"context"
	"prakarsa-app/transport/request"
	"prakarsa-app/transport/response"
)

type Partner struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    *bool  `json:"is_active"`
	CreatedBy   string `json:"created_by"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedBy   string `json:"updated_by"`
	UpdatedAt   int64  `json:"updated_at"`
	DeletedAt   int64  `json:"deleted_at"`
}

// // PartnerRepository represent the partner repository contract
type PartnerRepository interface {
	GetList(ctx context.Context, request *request.GetListPartnerReq) ([]response.GetListPartnerRes, response.MetaRes, error)
	Create(ctx context.Context, partner *Partner) error
	Update(ctx context.Context, partner *Partner) error
	Delete(ctx context.Context, partner *Partner) (int64, error)
}

// PartnerUsecase represent the partner usecase contract
type PartnerUsecase interface {
	GetList(ctx context.Context, request *request.GetListPartnerReq) ([]response.GetListPartnerRes, response.MetaRes, error)
	Create(ctx context.Context, request *request.CreatePartnerReq) (response.CreatePartnerRes, error)
	Update(ctx context.Context, request *request.UpdatePartnerReq) error
	Delete(ctx context.Context, request *request.DeletePartnerReq) (int64, error)
}
