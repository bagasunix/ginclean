package handlers

import (
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/bagasunix/ginclean/server/endpoints/middlewares"
	"github.com/gin-gonic/gin"
)

func MakeRoleHandler(eps endpoints.RoleEndpoint, rg *gin.RouterGroup) *gin.RouterGroup {
	rg.Use(middlewares.Auth())
	rg.Use(middlewares.Permission("guide", "user", "admin"))
	rg.POST("", eps.CreateRole())
	rg.POST("/create-multiple-role", eps.CreateMultiRole())
	rg.GET("", eps.ListRole())
	rg.GET("/:id", eps.ViewRole())
	rg.PATCH("/:id", eps.UpdateRole())
	rg.PATCH("/update-multiple-role", eps.UpdateMultiRole())
	rg.DELETE("/:id", eps.DeleteRole())
	return rg
}
