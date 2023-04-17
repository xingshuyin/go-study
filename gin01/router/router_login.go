package router

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	// "os"
	"time"
)

type token_ struct {
	Id   int64
	Name string
	jwt.StandardClaims
}

// var SECRET_KEY string = os.Getenv("SECRET_KEY")
var SECRET_KEY string = "asadasdadasdasdf"

func Token(id int64, name string) (signedToken string, signedRefresh string, err error) {
	claims := &token_{
		Id:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	claims_refresh := &token_{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	//创建token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
	}
	refresh, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims_refresh).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
	}
	return token, refresh, err
}

func ParseToken(token string) (t *token_, msg string) {
	//解析token
	token_data, err := jwt.ParseWithClaims(
		token,
		&token_{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		return
	}
	//获取token数据
	t, ok := token_data.Claims.(*token_)
	if !ok {
		msg = "token错误"
	}
	return
}
func Login(e *gin.RouterGroup) {
	e.POST("/login/", func(ctx *gin.Context) {
		username := ctx.PostForm("username") // 获取post请求的表单数据,不能获取post的json数据, json用ctx.BindJSON()
		password := ctx.PostForm("password")
		fmt.Println("username", username, password)
		if username != "" && password != "" {
			if token, refresh, err := Token(1, username); err == nil {
				ctx.JSON(200, gin.H{"token": token, "refresh": refresh})
			} else {
				ctx.AbortWithStatusJSON(401, gin.H{"detail": "账号或密码错误"})
			}
		} else {
			ctx.AbortWithStatusJSON(401, gin.H{"detail": "账号或密码为空"})
		}

	})
	e.GET("/login/", func(ctx *gin.Context) {
		username := ctx.Query("username")
		password := ctx.Query("password")
		if username != "" && password != "" {
			if token, refresh, err := Token(1, username); err == nil {
				ctx.JSON(200, gin.H{"token": token, "refresh": refresh})
			} else {
				ctx.AbortWithStatusJSON(401, gin.H{"detail": "账号或密码错误"})
			}
		} else {
			ctx.AbortWithStatusJSON(401, gin.H{"detail": "账号或密码错误"})
		}

	})
}
