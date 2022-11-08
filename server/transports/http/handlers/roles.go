package handlers

import (
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/gin-gonic/gin"
)

func MakeRoleHandler(eps endpoints.RoleEndpoint, rg *gin.RouterGroup) *gin.RouterGroup {
	rg.POST("", eps.CreateRole())
	rg.GET("", eps.ListRole())
	rg.GET("/:id", eps.ViewRole())
	rg.PATCH("/:id", eps.UpdateRole())
	rg.PATCH("/update-multiple-role", eps.UpdateMultiRole())
	rg.DELETE("/:id", eps.DeleteRole())
	return rg
}
