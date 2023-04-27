package main

import (
	"flag"
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
	r.AddFromFiles("json_format", "templates/base.html", "templates/json_format.html")
	return r
}

func main() {
	// 获取参数
	var port string
	var debug_flag bool
	flag.StringVar(&port, "port", "8080", "server listen port")
	flag.BoolVar(&debug_flag, "debug", true, "set debug or not")
	flag.Parse()

	// 设置启动模式
	if debug_flag {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.HTMLRender = createMyRender()
	// 首页
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{})
	})
	// 文本加解密
	router.GET("/text-encode-decode", func(c *gin.Context) {
		c.HTML(200, "text_encode_decode", gin.H{})
	})
	// 图片拼接
	router.GET("/stitch-picture", func(c *gin.Context) {
		c.HTML(200, "stitch_picture", gin.H{})
	})
	// 图片压缩
	router.GET("/compress-picture", func(c *gin.Context) {
		c.HTML(200, "compress_picture", gin.H{})
	})
	// JSON格式化
	router.GET("/json-format", func(c *gin.Context) {
		c.HTML(200, "json_format", gin.H{})
	})

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
