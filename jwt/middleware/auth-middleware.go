package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	helper "project/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("token") //TODO:获取请求头
		if clientToken == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("token is no0ne")})
			ctx.Abort()
			return
		}
		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}
		ctx.Set("email", claims.Email)
		ctx.Set("name", claims.Name)
	}
}
