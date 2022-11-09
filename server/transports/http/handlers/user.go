package handlers

import (
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/gin-gonic/gin"
)

func MakeUserHandler(eps endpoints.UserEndpoint, rg *gin.RouterGroup) *gin.RouterGroup {
	rg.POST("", eps.CreateUser())
	rg.GET("", eps.ListAccount())
	rg.GET("/:id", eps.ViewAccount())
	rg.DELETE("/:id", eps.DeleteAccount())
	rg.PUT("", eps.DisableAccount())
	rg.POST("/login", eps.LoginAccount())
	return rg
}
