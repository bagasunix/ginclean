package handlers

import (
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/bagasunix/ginclean/server/endpoints/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func MakeUserHandler(logs *zap.Logger, eps endpoints.UserEndpoint, rg *gin.RouterGroup) *gin.RouterGroup {
	rg.POST("", eps.CreateUser())
	rg.Use(middlewares.Auth(logs), middlewares.Permission("admin"))
	rg.GET("", eps.ListAccount())
	rg.GET("/:id", eps.ViewAccount())
	rg.DELETE("/:id", eps.DeleteAccount())
	rg.PUT("", eps.DisableAccount())
	return rg
}
