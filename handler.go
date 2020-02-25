package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func test(c *gin.Context) {
	c.String(http.StatusOK, "Your goMonitor is running!")
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", nil)
}

func showFiles(c *gin.Context) {
	var res string
	files, _ := ioutil.ReadDir("log")
	for _, file := range files {
		res += file.Name() + "\r\n"
	}
	c.String(http.StatusOK, res)
}

func showLog(c *gin.Context) {
	name := c.Query("name")
	if log, err := getLogfile(name); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.String(http.StatusOK, log)
	}
}

func parseLog(c *gin.Context) {
	filename := c.Query("name")
	from := c.Query("from")
	to := c.DefaultQuery("to", "@")
	if log, err := getLogfile(filename); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		res := parseFromTo(log, from, to)
		c.String(http.StatusOK, "Parse result from %s to %s is:%s", from, to, stringArray2string(res))
	}
}