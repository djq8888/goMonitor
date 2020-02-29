package main

import "strconv"

//字符串数组转为字符串
func stringArray2string(src []string) string {
	var res string
	for _, str := range src {
		res += "\r\n" + str
	}
	return res
}

//字符串数组转为int64数组
func stringArray2int64Array(src []string) []int64 {
	var res []int64
	for _, str := range src {
		tmp, _ := strconv.ParseInt(str, 10, 64)
		res = append(res, tmp)
	}
	return res
}

//int数组转为string数组
func intArray2stringArray(src []int) []string {
	var res []string
	for _, num := range src {
		tmp := strconv.Itoa(num)
		res = append(res, tmp)
	}
	return res
}
