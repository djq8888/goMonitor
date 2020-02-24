package main

func stringArray2string(src []string) string {
	var res string
	for _, str := range src {
		res += "\r\n" + str
	}
	return res
}
