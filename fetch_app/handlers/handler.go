package rest

import (
	"github.com/fetch_app/helper"
	"github.com/fetch_app/usecases/service"
	"github.com/gin-gonic/gin"
)

// Handler :nodoc:
type Handler struct {
	baseHandler
	service service.Service
}

// NewHandler :nodoc:
func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Fetch godoc
// @Summary Fetch Data
// @Tags Fetch App
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentication header"
// @Success 200 {object} []models.Resource "Ok"
// @Failure 500 {object} responses.Response
// @Router /fetch [get]
func (h *Handler) Fetch(ginCtx *gin.Context) {
	result, err := h.service.Fetch(ginCtx.Request.Context())

	if err != nil {
		h.errorResponse(ginCtx, err)
		return
	}

	h.successResponse(ginCtx, result)
}

// Aggregate godoc
// @Summary Aggregate area_province and date
// @Tags Fetch App
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentication header"
// @Success 200 {object} []models.Resource "Ok"
// @Failure 500 {object} responses.Response
// @Router /aggregate [get]
func (h *Handler) Aggregate(ginCtx *gin.Context) {
	result, err := h.service.Aggregate(ginCtx.Request.Context())

	if err != nil {
		h.errorResponse(ginCtx, err)
		return
	}

	h.successResponse(ginCtx, result)
}

// ClaimsJWT godoc
// @Summary claims jwt token
// @Tags Fetch App
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentication header"
// @Success 200 {object} helper.UserContext "Ok"
// @Failure 500 {object} responses.Response
// @Router /clamis-jwt [get]
func (h *Handler) ClaimsJWT(ginCtx *gin.Context) {
	userContext, err := helper.ParseUserContext(ginCtx.Request.Context())

	if err != nil {
		h.errorResponse(ginCtx, err)
		return
	}

	h.successResponse(ginCtx, userContext)
}
