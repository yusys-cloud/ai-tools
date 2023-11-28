// Author: yangzq80@gmail.com
// Date: 2023/8/17
// 字符串转换
package utils

import (
	"log"
	"strconv"
	"strings"
)

func StrToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("Convert %v to int error:%v", str, err.Error())
		return 0
	}
	return num
}

// 倒序数组
func ReverseArray(arr []string) []string {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		arr[i], arr[length-i-1] = arr[length-i-1], arr[i]
	}
	return arr
}

// 数组转换为字符
func ArrayToStr(arr []string) string {

	return strings.Join(arr, "\n")
}
