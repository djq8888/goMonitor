package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//打开日志文件并以string返回文件全部内容
func getLogfile(filename string) (string, error) {
	data, err := ioutil.ReadFile("log/"+filename)
	if err != nil {
		fmt.Println("File reading error: ", err)
		return "", err
	}
	return string(data), nil
}

//解析文本中所有from到to中间的字符串，并存入string数组
func parseFromTo(src, from, to string) []string {
	var res []string
	srcLen := len(src)
	destLen := len(from)
	start := 0
	for loc := 0; loc < srcLen;  {
		if loc = strings.Index(src[start:], from); loc > -1 {
			if endLoc := strings.Index(src[start + loc:], to); endLoc > -1 {
				value := src[start + loc + destLen : start + loc + endLoc]
				res = append(res, value)
				start += loc + endLoc
			}
		} else {
			break
		}
	}
	return res
}
