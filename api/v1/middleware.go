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
			if strings.Contains(ctx.Request.URL.Path, "api") {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				ctx.Abort()
			} else {
				ctx.HTML(http.StatusOK, "404.html", nil)
			}

			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		claims, err := internal.ValidateJWT(tokenString)

		if err != nil {
			if strings.Contains(ctx.Request.URL.Path, "api") {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				ctx.Abort()
			} else {
				ctx.HTML(http.StatusOK, "404.html", nil)
			}
			return
		}

		ctx.Set("claims", claims)

		ctx.Next()
	}
}
