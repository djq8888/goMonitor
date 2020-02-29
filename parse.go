package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unsafe"
)

//打开日志文件并以string返回文件全部内容
func getFile(path, filename string) (string, error) {
	data, err := ioutil.ReadFile(path+filename)
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

//解析文本中所有from到to中间的內容作为开始时间，from2:to2之间内容作为结束时间，并返回时间差
func parseInterval(src, from, to, from2, to2 string) []string {
	begin := parse2map(src, from, to)
	interval := findInterval(begin, src, from2, to2)
	return intArray2stringArray(interval)
}

//将时间戳序列，统计为qps，并存入string数组
func parseQps(data []string) []string {
	timestamp := stringArray2int64Array(data)
	start := timestamp[0]
	var qps []int
	i := 0
	qps = append(qps, 0)
	for _, ts := range timestamp {
		if ts > (start + 1000) {
			start += 1000
			qps = append(qps, 0)
			i++
		}
		qps[i]++
	}
	return intArray2stringArray(qps)
}

//解析監控文件中的CPU利用率，并存入string数组
func parseCPU(data string) []string {
	var res []string
	datas := strings.Split(data, "\n")
	for _, record := range datas {
		info := strings.Split(record, " ")
		res = append(res, info[0])
	}
	return res
}

//解析監控文件中的內存佔用率，并存入string数组
func parseMEM(data string) []string {
	var res []string
	datas := strings.Split(data, "\n")
	for _, record := range datas {
		info := strings.Split(record, " ")
		if len(info) > 1 {
			res = append(res, info[1])
		}
	}
	return res
}

//将from'value','key'to解析为map[key]value
func parse2map(src, from, to string) map[int]int64 {
	var resMap map[int]int64
	resMap = make(map[int]int64)
	srcLen := len(src)
	destLen := len(from)
	start := 0
	for loc := 0; loc < srcLen;  {
		if loc = strings.Index(src[start:], from); loc > -1 {
			if endLoc := strings.Index(src[start + loc:], to); endLoc > -1 {
				tmp := src[start + loc + destLen : start + loc + endLoc]
				res := strings.Split(tmp, ",")
				res_key, _ := strconv.Atoi(res[1])
				res_value, _ := strconv.ParseInt(res[0], 10, 64)
				resMap[res_key] = res_value
				start += loc + endLoc
			}
		} else {
			break
		}
	}
	return resMap
}

//将from'value','key'to解析为map[key]value，并与begin中对应key的value相减，将差值存为int数组
func findInterval(begin map[int]int64, src, from, to string) []int {
	var resArray []int
	srcLen := len(src)
	destLen := len(from)
	start := 0
	for loc := 0; loc < srcLen;  {
		if loc = strings.Index(src[start:], from); loc > -1 {
			if endLoc := strings.Index(src[start + loc:], to); endLoc > -1 {
				tmp := src[start + loc + destLen : start + loc + endLoc]
				res := strings.Split(tmp, ",")
				res_issi, _ := strconv.Atoi(res[1])
				res_time, _ := strconv.ParseInt(res[0], 10, 64)
				if v_begin, ok:= begin[res_issi]; ok {
					interval := res_time - v_begin
					if interval > 0 {
						//fmt.Println("error:", res_issi, v_begin, res_time)
						resArray = append(resArray, *(*int)(unsafe.Pointer(&interval)))
					}
				}
				start += loc + endLoc
			}
		} else {
			break
		}
	}
	return resArray
}