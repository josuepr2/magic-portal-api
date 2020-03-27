package server

import (
	"MagicPotal/server/middlewares"
	"github.com/gin-gonic/gin"
)

const (
	DEFAULTDB_CON = "defaultDB"
)

func Setup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.RedirectTrailingSlash = true

	r.Use(middlewares.DBInit(DEFAULTDB_CON))

	loadRoutes(r)

	return r
}
