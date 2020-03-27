package server

import (
	"MagicPotal/controllers/accounts"
	"MagicPotal/server/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func loadRoutes(r *gin.Engine) *gin.Engine {

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong ...")
	})

	core := r.Group("/core", middlewares.TXNInit(DEFAULTDB_CON))
	{
		core.POST("/accounts", accounts.CreateAccount)
		core.PUT("/accounts/:aid", accounts.UpdateAccount)
		core.GET("/accounts", accounts.GetAccounts)
		core.GET("/accounts/:aid", accounts.GetSingleAccount)

	}

	return r
}