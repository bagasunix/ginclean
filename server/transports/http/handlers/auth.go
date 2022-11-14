package handlers

import (
	"github.com/bagasunix/ginclean/server/endpoints"
	"github.com/gin-gonic/gin"
)

func MakeAuthHandler(eps endpoints.AuthEndpoint, rg *gin.RouterGroup) *gin.RouterGroup {
	rg.POST("/refresh_token", eps.RefreshToken())
	rg.POST("/login", eps.LoginAccount())
	return rg
}
