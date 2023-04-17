package routes

import (
	controlers "project/controlers"

	"github.com/gin-gonic/gin"
)

func Auth_router(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controlers.Signup())
	incomingRoutes.POST("users/login", controlers.Login())
}
