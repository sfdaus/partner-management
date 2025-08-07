package usecase

import (
	"context"
	"fmt"
	"prakarsa-app/transport/response"
	"time"

	"prakarsa-app/domain"
	"prakarsa-app/repository/redis"
	"prakarsa-app/transport/request"

	"github.com/google/uuid"
)

type PartnerUsecase struct {
	partnerRepo domain.PartnerRepository
	redisRepo   redis.RedisRepository
	ctxTimeout  time.Duration
}

// NewPartnerUsecase
func NewPartnerUsecase(partnerRepo domain.PartnerRepository, redisRepo redis.RedisRepository, ctxTimeout time.Duration) *PartnerUsecase {
	return &PartnerUsecase{
		partnerRepo: partnerRepo,
		redisRepo:   redisRepo,
		ctxTimeout:  ctxTimeout,
	}
}
func (u *PartnerUsecase) GetList(c context.Context, request *request.GetListPartnerReq) (res []response.GetListPartnerRes, meta response.MetaRes, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	res, meta, err = u.partnerRepo.GetList(ctx, request)

	return
}
func (u *PartnerUsecase) GetDetail(c context.Context, request *request.GetDetailPartnerReq) (res domain.Partner, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	res, err = u.partnerRepo.GetDetail(ctx, request)

	return
}
func (u *PartnerUsecase) Create(c context.Context, request *request.CreatePartnerReq) (res response.CreatePartnerRes, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	// Create Payload
	partnerID := uuid.NewString()
	t := true
	partnerPayload := &domain.Partner{
		ID:          partnerID,
		Name:        request.Name,
		Description: request.Description,
		IsActive:    &t,
		CreatedBy:   "TODO_created_by",
		CreatedAt:   time.Now().Unix(),
	}

	// Response Payload
	res.ID = partnerID

	err = u.partnerRepo.Create(ctx, partnerPayload)
	return
}
func (u *PartnerUsecase) Update(c context.Context, request *request.UpdatePartnerReq) (err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	// Update Payload
	partnerPayload := &domain.Partner{
		ID:        request.ID,
		UpdatedBy: "TODO_updated_by",
		UpdatedAt: time.Now().Unix(),
	}

	if request.Name != "" {
		partnerPayload.Name = request.Name
	}
	fmt.Println(request)
	if request.Description != "" {
		partnerPayload.Description = request.Description
	}

	if request.IsActive != nil {
		partnerPayload.IsActive = request.IsActive
	}

	err = u.partnerRepo.Update(ctx, partnerPayload)
	return
}
func (u *PartnerUsecase) Delete(c context.Context, request *request.DeletePartnerReq) (rowsAffected int64, err error) {
	ctx, cancel := context.WithTimeout(c, u.ctxTimeout)
	defer cancel()

	threadPayload := &domain.Partner{
		ID: request.ID,
	}

	rowsAffected, err = u.partnerRepo.Delete(ctx, threadPayload)
	return
}
