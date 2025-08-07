package http

import (
	"net/http"
	"prakarsa-app/delivery/middleware"
	"prakarsa-app/domain"
	"prakarsa-app/transport/request"
	"prakarsa-app/utils"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

type PartnerHandler struct {
	PartnerUC domain.PartnerUsecase
}

// NewPartnerHandler will initialize the todo resources endpoint
func NewPartnerHandler(e *echo.Echo, middleware *middleware.Middleware, partnerUC domain.PartnerUsecase) {
	handler := &PartnerHandler{
		PartnerUC: partnerUC,
	}

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/partner-types", handler.GetList)
	apiV1.POST("/partner-types", handler.Create)
	apiV1.PATCH("/partner-types/:id", handler.Update)
	apiV1.DELETE("/partner-types/:id", handler.Delete)
}

func (h *PartnerHandler) GetList(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.GetListPartnerReq

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewUnprocessableEntityError(err.Error()))
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, utils.NewInvalidInputError(errVal))
	}

	if res, meta, err := h.PartnerUC.GetList(ctx, &req); err != nil {
		return c.JSON(utils.ParseHttpError(err))
	} else {
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Partner Type successfully retrieved",
			"data":    res,
			"meta":    meta,
		})
	}

}

func (h *PartnerHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.CreatePartnerReq

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewUnprocessableEntityError(err.Error()))
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, utils.NewInvalidInputError(errVal))
	}

	if res, err := h.PartnerUC.Create(ctx, &req); err != nil {
		return c.JSON(utils.ParseHttpError(err))
	} else {
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Partner Type successfully created",
			"data":    res,
		})
	}

}

func (h *PartnerHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.UpdatePartnerReq

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewUnprocessableEntityError(err.Error()))
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, utils.NewInvalidInputError(errVal))
	}

	if err := h.PartnerUC.Update(ctx, &req); err != nil {
		return c.JSON(utils.ParseHttpError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Partner Type successfully updated",
	})
}

func (h *PartnerHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.DeletePartnerReq

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewUnprocessableEntityError(err.Error()))
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, utils.NewInvalidInputError(errVal))
	}

	if rowsAffected, err := h.PartnerUC.Delete(ctx, &req); err != nil {
		return c.JSON(utils.ParseHttpError(err))
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Partner Type successfully deleted",
			"data": map[string]int64{
				"rows_affected": rowsAffected,
			},
		})
	}
}
