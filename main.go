package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

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
