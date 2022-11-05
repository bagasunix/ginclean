package handlers

import (
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/bagasunix/ginclean/server/endpoints/middlewares"
	"github.com/gin-gonic/gin"
)

func MakeUserHandler(eps endpoints.RoleEndpoint, rg *gin.RouterGroup) *gin.RouterGroup {
	rg.Use(middlewares.CORSMiddleware())
	rg.POST("", eps.CreateRole())
	rg.GET("", eps.ListRole())
	rg.GET("/:id", eps.ViewRole())
	rg.PATCH("/:id", eps.UpdateRole())
	rg.PATCH("/update-multiple-role", eps.UpdateMultiRole())
	rg.DELETE("/:id", eps.DeleteRole())
	return rg
}
