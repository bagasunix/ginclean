package handlers

import (
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/bagasunix/ginclean/server/endpoints/middlewares"
	"github.com/gin-gonic/gin"
)

func MakeUserHandler(eps endpoints.UserEndpoint, rg *gin.RouterGroup) *gin.RouterGroup {
	rg.POST("", eps.CreateUser())
	rg.POST("/login", eps.LoginAccount())
	rg.GET("", middlewares.Auth(), middlewares.Permission("admin"), eps.ListAccount())
	rg.GET("/:id", middlewares.Auth(), middlewares.Permission("admin"), eps.ViewAccount())
	rg.DELETE("/:id", middlewares.Auth(), middlewares.Permission("admin"), eps.DeleteAccount())
	rg.PUT("", middlewares.Auth(), middlewares.Permission("admin"), eps.DisableAccount())
	return rg
}
