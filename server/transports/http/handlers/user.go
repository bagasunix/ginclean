package handlers

import (
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/bagasunix/ginclean/server/endpoints/middlewares"
	"github.com/gin-gonic/gin"
)

func MakeUserHandler(eps endpoints.UserEndpoint, rg *gin.RouterGroup) *gin.RouterGroup {
	rg.Use(middlewares.CORSMiddleware())
	rg.POST("", eps.CreateAccount())
	return rg
}
