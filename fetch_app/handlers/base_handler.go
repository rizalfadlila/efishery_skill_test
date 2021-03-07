package rest

import (
	"fmt"
	"net/http"

	"github.com/fetch_app/constants"
	responses "github.com/fetch_app/handlers/response"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type baseHandler struct {
}

// SuccessResponse :nodoc:
func (handler *baseHandler) successResponse(ginCtx *gin.Context, data interface{}) {
	ginCtx.JSON(http.StatusOK, responses.Response{
		Data:    data,
		Message: "Successfully Retrieve Data",
		Status:  constants.StatusSuccess,
	})
}

// ErrorResponse :nodoc:
func (handler *baseHandler) errorResponse(ginCtx *gin.Context, err error) {
	ginCtx.JSON(http.StatusInternalServerError, responses.Response{
		Errors: parseError(err),
		Status: constants.StatusFailed,
	})
}

// parseError :nodoc:
func parseError(err error) []string {
	if err == nil {
		return nil
	}

	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return []string{err.Error()}
	}

	var errors []string
	for _, e := range ve {
		errors = append(errors, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", e.Field(), e.Tag()))
	}

	return errors
}
