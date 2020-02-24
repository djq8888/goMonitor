package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Your goMonitor is running!")
	})
	r.GET("/showLog", func(c *gin.Context) {
		if log, err := getLogfile("test.log"); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.String(http.StatusOK, log)
		}
	})
	r.GET("/showFiles", func(c *gin.Context) {
		res := showFiles(".")
		c.String(http.StatusOK, res)
	})
	r.GET("/parseLog", func(c *gin.Context) {
		from := c.Query("from")
		to := c.DefaultQuery("to", "@")
		if log, err := getLogfile("test.log"); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			res := parseFromTo(log, from, to)
			c.String(http.StatusOK, "Parse result from %s to %s is:%s", from, to, stringArray2string(res))
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
