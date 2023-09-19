package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, nil)
	})

	return nil
}
