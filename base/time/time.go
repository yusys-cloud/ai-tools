// Author: yangzq80@gmail.com
// Date: 2023/4/27
package time

import "time"

const dateTimeFormat = "2006-01-02 15:04:05"
const dateFormat = "2006-01-02"

func GetCurrentDateTime() string {
	return time.Now().Format(dateTimeFormat)
}

func GetCurrentDate() string {
	return time.Now().Format(dateFormat)
}
