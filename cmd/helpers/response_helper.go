package helpers

import (
	"net/http"
	"service-fleetime/cmd/schemes"

	"github.com/gin-gonic/gin"
)

func APIResponsePaginated(
	ctx *gin.Context,
	Message string,
	StatusCode int,
	Data interface{},
	Count int64,
	Page int,
	TotalPage int,
) {

	jsonResponse := schemes.SchemeResponsePaginated{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
		Count:      Count,
		Page:       Page,
		TotalPage:  TotalPage,
	}

	if StatusCode >= 400 {
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Data interface{}) {

	jsonResponse := schemes.SchemeResponses{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}

}

func ErrorResponse(ctx *gin.Context, Error interface{}) {
	err := schemes.SchemeErrorResponse{
		StatusCode: http.StatusBadRequest,
		Error:      Error,
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
}
