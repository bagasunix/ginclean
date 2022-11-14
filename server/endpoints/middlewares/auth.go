package middlewares

import (
	"fmt"
	"strings"

	"github.com/bagasunix/ginclean/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func Auth(logs *zap.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader(authorizationHeaderKey)
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "You are not logged in! Please log in to get access."})
			context.Abort()
			return
		}

		fields := strings.Fields(tokenString)
		if len(fields) < 2 {
			context.JSON(401, gin.H{"error": "invalid authorization header format"})
			context.Abort()
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			context.JSON(401, gin.H{"error": fmt.Errorf("unsupported authorization type %s", authorizationType)})
			context.Abort()
			return
		}

		claims, err := jwt.ValidateToken(logs, fields[1])
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Set(authorizationPayloadKey, *claims.User)
		context.Next()
	}
}
