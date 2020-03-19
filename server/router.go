package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loadRoutes(r *gin.Engine) *gin.Engine {

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong ...")
	})

	return r
}