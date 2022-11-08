package handlers

import (
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/gin-gonic/gin"
)

func MakeUserHandler(eps endpoints.UserEndpoint, rg *gin.RouterGroup) *gin.RouterGroup {
	rg.POST("", eps.CreateUser())
	rg.GET("", eps.ListAccount())
	rg.DELETE("/:id", eps.DeleteAccount())
	return rg
}
