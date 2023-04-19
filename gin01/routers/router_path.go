package routers

import (
	"fmt"
	models "project/models"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type interfaces struct {
	method string
	path   string
}

func Logging(ctx *gin.Context) {
	time.Sleep(time.Second)
	fmt.Println("Logging", ctx.Request.URL)
}

func Permission(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token") //TODO:获取请求头
	if token_data, msg := ParseToken(token); msg == "" && token != "" {
		fmt.Println("token", token, token_data)
		go Logging(ctx.Copy()) // 使用协程, ctx.Copy()使用此时的上下文,不会被后续中间件影响
		method := ctx.Request.Method
		path := ctx.Request.URL
		fmt.Println(path)
		rules := []interfaces{
			{method: "GET", path: "/user/.*?/"},
			{method: "PUT", path: "/user/.*?/"},
			{method: "GET", path: "/user/"},
			{method: "POST", path: "/user/"},
			{method: "GET", path: "/menu/"}}
		is_found := false
		for _, rule := range rules {
			if found, err := regexp.MatchString(rule.path, path.String()+"\\??.*"); err == nil {
				if strings.EqualFold(rule.method, method) && found {
					is_found = found
					break
					// ctx.Next() // 执行后面所有的和中间件
				}
			}
		}
		if !is_found {
			ctx.AbortWithStatusJSON(401, gin.H{"detail": "没有权限"}) // 阻止后续中间件,但会执行后面的代码
			return
		}
	} else {
		ctx.AbortWithStatusJSON(401, gin.H{"detail": msg})
	}
}

type ModelType interface {
	models.User | models.Menu
}

type ModelList[T ModelType] []T

// RESTFUL路由
func Path[T ModelType](e *gin.RouterGroup, base_url string, model_example T) {
	e.GET(fmt.Sprintf("/%s/", base_url), Permission, func(ctx *gin.Context) {
		model_type := reflect.TypeOf(model_example)
		numfield := model_type.NumField()
		fmt.Println("model_type", model_type, numfield, reflect.TypeOf(numfield))
		var fields = make([]string, numfield)
		for i := 0; i < numfield; i++ {
			fields[i] = model_type.Field(i).Name
		}
		fmt.Println(fields)
		page, err := strconv.ParseInt(ctx.Query("page"), 10, 64)
		if err != nil {
			fmt.Println("err", err)
			ctx.JSON(400, gin.H{"detail": "未分页"})
			return
		}
		limit, err := strconv.ParseInt(ctx.Query("limit"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"detail": "未限制数量"})
			return
		}
		var filter T
		ctx.ShouldBind(&filter)
		fmt.Println("filter", filter)
		var items ModelList[T]
		models.DB.Where(filter).Limit(int(limit)).Offset(int(limit) * (int(page) - 1)).Find(&items)
		ctx.JSONP(200, items)
	})
	e.GET(fmt.Sprintf("/%s/:id/", base_url), Permission, func(ctx *gin.Context) {
		id, ok := ctx.Params.Get("id")
		if ok {
			id_int, err := strconv.ParseInt(id, 10, 32)
			if err != nil {
				ctx.JSONP(404, gin.H{"detail": "id错误"})
			} else {
				var item T
				models.DB.First(&item, id_int)
				ctx.JSONP(200, item)
			}
		} else {
			ctx.JSONP(404, gin.H{"detail": "请求错误"})
		}
	})
	e.PUT(fmt.Sprintf("/%s/:id/", base_url), Permission, func(ctx *gin.Context) {
		if id, ok := ctx.Params.Get("id"); ok { // 获取路径参数
			data := make(map[string]interface{})        //注意该结构接受的内容
			if err := ctx.BindJSON(&data); err == nil { //获取post的json数据
				if id_int, err := strconv.ParseInt(id, 10, 32); err != nil {
					ctx.JSONP(404, gin.H{"detail": "id错误"})
				} else {
					var item T
					models.DB.First(&item, id_int)
					models.DB.Model(&item).Updates(data)
					ctx.JSONP(200, gin.H{"detail": "更新成功", "data": item})
				}
			}
		} else {
			ctx.JSONP(404, gin.H{"detail": "请求错误"})
		}

	})
	e.POST(fmt.Sprintf("/%s/", base_url), Permission, func(ctx *gin.Context) {
		var item T
		ctx.BindJSON(&item) //获取post的json数据
		r := models.DB.Create(&item)
		if r.Error == nil {
			ctx.JSONP(200, gin.H{"detail": "创建成功", "data": item})
		} else {
			ctx.JSON(400, gin.H{"detail": "已存在"})
		}
	})
}
