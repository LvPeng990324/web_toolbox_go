package main

import (
	"log"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/base.html", "templates/index.html")
	r.AddFromFiles("text_encode_decode", "templates/base.html", "templates/text_encode_decode.html")
	r.AddFromFiles("stitch_picture", "templates/base.html", "templates/stitch_picture.html")
	r.AddFromFiles("compress_picture", "templates/base.html", "templates/compress_picture.html")
	return r
}

func main() {
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.HTMLRender = createMyRender()
	// 首页
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			
		})
	})
	// 文本加解密
	router.GET("/text-encode-decode", func(c *gin.Context) {
		c.HTML(200, "text_encode_decode", gin.H{

		})
	})
	// 图片拼接
	router.GET("/stitch-picture", func(c *gin.Context) {
		c.HTML(200, "stitch_picture", gin.H{

		})
	})
	// 图片压缩
	router.GET("/compress-picture", func(c *gin.Context) {
		c.HTML(200, "compress_picture", gin.H{

		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}