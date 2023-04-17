package router

import (
	"fmt"
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

func middleware(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token") //TODO:获取请求头
	if token_data, msg := ParseToken(token); msg == "" && token != "" {
		fmt.Println("token", token, token_data, token_data.Id)
		go Logging(ctx.Copy()) // 使用协程, ctx.Copy()使用此时的上下文,不会被后续中间件影响
		method := ctx.Request.Method
		path := ctx.Request.URL
		fmt.Println(path)
		rules := []interfaces{{method: "GET", path: "/user/.*?/"}, {method: "PUT", path: "/user/.*?/"}, {method: "GET", path: "/user/"}}
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
		ctx.AbortWithStatusJSON(401, gin.H{"detail": "无效令牌"})
	}
}

func Path(e *gin.RouterGroup, base_url string, model interface{}) {
	model_type := reflect.TypeOf(model)
	model_value := reflect.ValueOf(model)
	fmt.Println(model_type, model_value)
	e.GET(fmt.Sprintf("/%s/", base_url), middleware, func(ctx *gin.Context) {
		// query := ctx.Request.URL.Query()
		// fmt.Println(query)
		ctx.JSONP(200, []map[string]interface{}{{"id": "1"}, {"id": "2"}})
	})
	e.GET(fmt.Sprintf("/%s/:id/", base_url), middleware, func(ctx *gin.Context) {
		id, ok := ctx.Params.Get("id")
		if ok {
			id_int, err := strconv.ParseInt(id, 10, 32)
			if err != nil {
				ctx.JSONP(404, map[string]interface{}{"detail": "id错误"})
			} else {
				ctx.JSONP(200, map[string]interface{}{"id": id_int, "GET": true})
			}
		} else {
			ctx.JSONP(404, map[string]interface{}{"detail": "请求错误"})
		}
	})
	e.PUT(fmt.Sprintf("/%s/:id/", base_url), middleware, func(ctx *gin.Context) {
		id, ok := ctx.Params.Get("id")       // 获取路径参数
		data := make(map[string]interface{}) //注意该结构接受的内容
		ctx.BindJSON(&data)                  //获取post的json数据
		fmt.Println("data", data)
		if ok {
			id_int, err := strconv.ParseInt(id, 10, 32)
			if err != nil {
				ctx.JSONP(404, map[string]interface{}{"detail": "id错误"})
			} else {
				ctx.JSONP(200, map[string]interface{}{"id": id_int})
			}
		} else {
			ctx.JSONP(404, map[string]interface{}{"detail": "请求错误"})
		}
	})
	e.POST(fmt.Sprintf("/%s/", base_url), middleware, func(ctx *gin.Context) {
		data := make(map[string]interface{}) //注意该结构接受的内容
		ctx.BindJSON(&data)                  //获取post的json数据
		fmt.Println("data", data)
		ctx.JSONP(200, map[string]interface{}{"id": "new_id"})
	})
}
