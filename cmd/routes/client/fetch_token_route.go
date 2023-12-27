package client

import (
	"service-fleetime/cmd/controller"

	"github.com/gin-gonic/gin"
)

func FetchToken(r *gin.Engine) {
	r.POST("/token-client", controller.GetToken)
}
