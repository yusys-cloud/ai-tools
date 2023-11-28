// Author: yangzq80@gmail.com
// Date: 2023/4/27
package utils

import "time"

const DateTimeFormat = "2006-01-02 15:04:05"
const DateFormat = "2006-01-02"

func GetCurrentDateTime() string {
	return time.Now().Format(DateTimeFormat)
}

func GetCurrentDate() string {
	return time.Now().Format(DateFormat)
}
