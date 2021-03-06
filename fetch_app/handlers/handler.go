package rest

import (
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
// @Success 200 {object} []models.Fetch "Ok"
// @Failure 500 {object} responses.Response
// @Router /fetch [get]
func (h *Handler) Fetch(ginCtx *gin.Context) {
	result, err := h.service.Fetch(ginCtx.Request.Context())

	if err != nil {
		h.errorResponse(ginCtx)
		return
	}

	h.successResponse(ginCtx, result)
}
