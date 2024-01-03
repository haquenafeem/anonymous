package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/haquenafeem/anonymous/internal"
)

func (api *Api) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()

			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		claims, err := internal.ValidateJWT(tokenString)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()

			return
		}

		ctx.Set("claims", claims)

		ctx.Next()
	}
}
