package routes

import (
	controlers "project/controlers"
	"project/middleware"

	"github.com/gin-gonic/gin"
)

func User_router(r *gin.Engine) {
	r.Use(middleware.Authenticate()) //  .Use(middleware.Authenticate())
	r.GET("/users", controlers.Get_users())
	r.GET("/users/:user_id", controlers.Get_user())
}
