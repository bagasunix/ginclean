package handlers

import (
	"github.com/bagasunix/ginclean/endpoints"
	"github.com/gin-gonic/gin"
)

func MakeUserHandler(eps endpoints.RoleEndpoint, rg *gin.RouterGroup) *gin.RouterGroup {
	rg.POST("", eps.CreateRole())
	return rg
}