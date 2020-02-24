package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//加载html模板
	r.LoadHTMLGlob("templates/*")

	//html模板调用的js
	r.Static("/js", "js")

	//主页
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", nil)
	})

	//测试接口
	r.GET("/test", test)
	//展示路径下的所有文件
	r.GET("/showFiles", showFiles)
	//展示日志全部内容
	r.GET("/showLog", showLog)
	//展示日志解析后的内容
	r.GET("/parseLog", parseLog)

	r.Run() // listen and serve on 0.0.0.0:8080
}
