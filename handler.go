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

func analyse(c *gin.Context) {
	c.HTML(http.StatusOK, "analyse.tmpl", nil)
}

func monitor(c *gin.Context) {
	c.HTML(http.StatusOK, "monitor.tmpl", nil)
}

func showLogs(c *gin.Context) {
	var res string
	files, _ := ioutil.ReadDir("log")
	for _, file := range files {
		res += file.Name() + "\r\n"
	}
	c.String(http.StatusOK, res)
}

func showMonitors(c *gin.Context) {
	var res string
	files, _ := ioutil.ReadDir("processMonitor")
	for _, file := range files {
		res += file.Name() + "\r\n"
	}
	c.String(http.StatusOK, res)
}

func showLog(c *gin.Context) {
	name := c.Query("name")
	if log, err := getFile("log/", name); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.String(http.StatusOK, log)
	}
}

func parseLog(c *gin.Context) {
	filename := c.Query("name")
	from := c.Query("from")
	to := c.DefaultQuery("to", "@")
	from2 := c.DefaultQuery("from2", "@")
	to2 := c.DefaultQuery("to2", "@")
	if log, err := getFile("log/", filename); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		var res []string
		if from2 == "@" {
			res = parseFromTo(log, from, to)
		} else {
			res = parseInterval(log, from, to, from2, to2)
		}
		c.String(http.StatusOK, "Parse result from %s to %s is:%s", from, to, stringArray2string(res))
	}
}

func getQps(c *gin.Context) {
	filename := c.Query("name")
	from := c.Query("from")
	to := c.DefaultQuery("to", "@")
	if log, err := getFile("log/", filename); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		res := parseFromTo(log, from, to)
		qps := parseQps(res)
		c.String(http.StatusOK, "Qps from %s to %s is:%s", from, to, stringArray2string(qps))
	}
}

func getCPU(c *gin.Context) {
	filename := c.Query("name")
	if data, err := getFile("processMonitor/", filename); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		res := parseCPU(data)
		c.String(http.StatusOK, stringArray2string(res))
	}
}

func getMEM(c *gin.Context) {
	filename := c.Query("name")
	if data, err := getFile("processMonitor/", filename); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		res := parseMEM(data)
		c.String(http.StatusOK, stringArray2string(res))
	}
}