package notfound

import (
	"net/http"
	"service-fleetime/cmd/helpers"

	"github.com/gin-gonic/gin"
)

func Roadblock(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		helpers.APIResponse(c, "What the hell are you doing", http.StatusNotFound, nil)
	})
}
