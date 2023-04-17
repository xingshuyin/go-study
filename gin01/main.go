package main

import (
	"fmt"
	models "project/model"
	"project/router"

	"github.com/gin-gonic/gin"
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
	router.Login(admin)
	router.Upload(admin)
	router.Path(admin, "user", map[string]string{})
	router.Path(admin, "menu", map[string]string{})
	// fmt.Println(e.Routes()) //获取所有路由
	e.Run()
}
