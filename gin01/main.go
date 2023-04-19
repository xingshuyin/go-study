package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	models "project/models"
	"project/routers"
	_ "project/tools"
)

func middleware_global(ctx *gin.Context) {
	fmt.Println("middleware_global")
}

func middleware_group(ctx *gin.Context) {
	fmt.Println("middleware_group")
}

func main() {
	e := gin.Default()
	models.Init()
	e.Static("./static", "./static") //配置静态目录
	e.Use(middleware_global)
	admin := e.Group("/api")
	admin.Use(middleware_group) // 使用全局中间件
	routers.Login(admin)
	routers.Upload(admin)
	routers.Path(admin, "user", models.User{})
	routers.Path(admin, "menu", models.Menu{})
	// router.Path(admin, "menu", map[string]string{})
	// fmt.Println(e.Routes()) //获取所有路由
	e.Run()
}
