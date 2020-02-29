package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//加载html模板
	r.LoadHTMLGlob("templates/*")
	//html模板调用的js
	r.Static("/js", "js")

	//测试接口
	r.GET("/test", test)
	//主页
	r.GET("/home", home)
	//日志分析平台
	r.GET("/analyse", analyse)
	//性能监控平台
	r.GET("/monitor", monitor)
	//获取所有日志
	r.GET("/showLogs", showLogs)
	//获取所有监控文件
	r.GET("/showMonitors", showMonitors)
	//获取日志全部内容
	r.GET("/showLog", showLog)
	//获取日志解析后的内容
	r.GET("/parseLog", parseLog)
	//获取Qps信息
	r.GET("/getQps", getQps)
	//获取CPU信息
	r.GET("/getCPU", getCPU)
	//获取內存信息
	r.GET("/getMEM", getMEM)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
