package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/base.html", "templates/index.html")
	r.AddFromFiles("text_encode_decode", "templates/base.html", "templates/text_encode_decode.html")
	r.AddFromFiles("stitch_picture", "templates/base.html", "templates/stitch_picture.html")
	r.AddFromFiles("compress_picture", "templates/base.html", "templates/compress_picture.html")
	r.AddFromFiles("json_format", "templates/base.html", "templates/json_format.html")
	r.AddFromFiles("timestamp_covert", "templates/base.html", "templates/timestamp_covert.html")
	r.AddFromFiles("chatgpt", "templates/base.html", "templates/chatgpt.html")
	r.AddFromFiles("port_scan", "templates/base.html", "templates/port_scan.html")
	r.AddFromFiles("markdown_edit", "templates/base.html", "templates/markdown_edit.html")
	r.AddFromFiles("crontab_test", "templates/base.html", "templates/crontab_test.html")
	r.AddFromFiles("fast_copy", "templates/base.html", "templates/fast_copy.html")
	return r
}

func main() {
	// 获取配置
	viper.SetConfigName("conf")
	viper.SetConfigType("ini")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	port := viper.GetString("system.port")
	debug_flag := viper.GetBool("system.debug")
	chatgpt_server_url := viper.GetString("chatgpt.server_url")
	chatgpt_api_key := viper.GetString("chatgpt.api_key")

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
	// 时间戳转换
	router.GET("/timestamp-covert", func(c *gin.Context) {
		c.HTML(200, "timestamp_covert", gin.H{})
	})
	// ChatGPT
	router.GET("/chatgpt", func(c *gin.Context) {
		c.HTML(200, "chatgpt", gin.H{})
	})

	// 获取ChatGPT回答
	router.POST("/chatgpt/ask", func(c *gin.Context) {
		check_res := check_req_valid(c)
		if !check_res { // 验证不过
			c.JSON(403, gin.H{"err": "permission denied!"})
			return
		}

		ask_content := c.PostForm("ask_content")
		target_url := chatgpt_server_url

		// 拼接参数
		payload := url.Values{}
		payload.Set("api_key", chatgpt_api_key)
		payload.Set("ask_content", ask_content)

		// 发送请求
		request, err := http.NewRequest(http.MethodPost, target_url, strings.NewReader(payload.Encode()))
		if err != nil {
			log.Fatal(err)
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()

		// 解析body
		body, _ := ioutil.ReadAll(response.Body)
		body_str := string(body)

		// 打个日志，记录下都问了啥
		log.Printf("chatgpt, ask: " + ask_content + " answer: " + body_str)

		c.JSON(200, gin.H{
			"answer": body_str,
		})
	})

	// 端口扫描
	router.GET("/port-scan", func(c *gin.Context) {
		c.HTML(200, "port_scan", gin.H{})
	})

	// markdown编辑
	router.GET("/markdown-edit", func(c *gin.Context) {
		c.HTML(200, "markdown_edit", gin.H{})
	})

	// crontab测试
	router.GET("/crontab-test", func(c *gin.Context) {
		c.HTML(200, "crontab_test", gin.H{})
	})

	// 快捷粘贴板
	router.GET("/fast-copy", func(c *gin.Context) {
		c.HTML(200, "fast_copy", gin.H{})
	})

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

// 验证请求合法性，就先简单的秒级时间戳乘以2吧
func check_req_valid(c *gin.Context) bool {
	// 获取参数里带的
	token := c.GetHeader("req_token")
	token_int, _ := strconv.Atoi(token)
	token_timestamp := token_int / 2

	timestamp := int(time.Now().Unix())

	// 给5s的时间容错
	return timestamp-token_timestamp <= 5
}
