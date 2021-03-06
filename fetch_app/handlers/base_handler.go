package rest

import (
	"net/http"

	"github.com/fetch_app/constants"
	responses "github.com/fetch_app/handlers/response"

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
func (handler *baseHandler) errorResponse(ginCtx *gin.Context) {
	ginCtx.JSON(http.StatusInternalServerError, responses.Response{
		Message: "Something went wrong",
		Status:  constants.StatusFailed,
	})
}
