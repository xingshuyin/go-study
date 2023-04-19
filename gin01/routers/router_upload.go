package routers

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

type File struct {
	Id   int64
	Url  string
	Name string
}

func FilePath() {

}

func Upload(e *gin.RouterGroup) {
	e.POST("/upload/", func(ctx *gin.Context) { //单上传文件
		if file, err := ctx.FormFile("file"); err == nil {
			fmt.Println(file)
			file_path := path.Join(time.Now().Format("./upload/2006/01/02/"), fmt.Sprint(time.Now().Unix())+file.Filename) //所有项目的路径都是根据main.go文件的相对路径
			ctx.SaveUploadedFile(file, file_path)
			ctx.JSON(200, File{Id: time.Now().Unix(), Url: file_path, Name: file.Filename})
		} else {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(400, gin.H{"detail": "上传失败"})
		}
	})
	e.POST("/uploads/", func(ctx *gin.Context) { //多文件上传
		uploaded := []File{}
		if form, err := ctx.MultipartForm(); err == nil {
			files := form.File["file"]
			for _, file := range files {
				file_name := file.Filename
				// file_size := file.Size
				file_path := path.Join(time.Now().Format("./upload/2006/01/02/"), fmt.Sprint(time.Now().Unix())+file_name)
				err := ctx.SaveUploadedFile(file, file_path)
				if err != nil {
					fmt.Println("上传失败------------")
					for _, v := range uploaded {
						os.Remove(v.Url)
					}
					ctx.AbortWithStatusJSON(400, gin.H{"detail": "上传失败"})
					return
				} else {
					uploaded = append(uploaded, File{Id: time.Now().Unix(), Url: file_path, Name: file_name})
				}
			}
			fmt.Println("uploaded", uploaded)
			ctx.JSON(200, uploaded)
		}
	})
}
