package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"path"
	_ "task5/docs"
	JWT "task5/jwt"
)

//@title test

//@version 1.0

//@description task5

func main() {
	r := gin.Default()
	r.GET("/hello", Hello)
	r.POST("/upload", Upload)
	r.POST("/download", Download)
	//获取token
	r.GET("/token", Token)
	//测试token
	r.GET("/test", JWT.JwtAuth(), Test)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run()
}

// @Summary 测试用token鉴权
// @Success 200 {object} gin.H
// @Router /test [get]
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// @Summary 获取token
// @Success 200 {object} gin.H
// @Router /token [get]
func Token(c *gin.Context) {
	JWT.GenerateToken(c)
}

// @Summary 下载文件
// @Success 200 {object} gin.H
// @Router /download [post]
func Download(c *gin.Context) {
	f := c.PostForm("file")
	c.FileAttachment(path.Join("./", f), f)
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// @Summary 上传文件
// @Success 200 {object} gin.H
// @Router /upload [post]
func Upload(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	dst := path.Join("./", f.Filename)
	err = c.SaveUploadedFile(f, dst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// @Summary 返回hello gin
// @Success 200 {string} string    "Hello gin"
// @Router /hello [get]
func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello gin")
}
