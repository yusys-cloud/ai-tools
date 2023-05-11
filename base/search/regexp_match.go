// Author: yangzq80@gmail.com
// Date: 2021/10/6
package search

import (
	"regexp"
)

func ContainsStr(str string, strs []string) bool {
	for _, e := range strs {
		if regexp.MustCompile(e).MatchString(str) {
			//if strings.Contains(content, e) {
			return true
		}
	}
	return false
}

// 通过正则从字符串中提取匹配的唯一值
func ExtractString(str string, regexpStr string) string {
	matches := ExtractMultipleString(str, regexpStr)
	if len(matches) > 1 {
		return matches[1]
	} else {
		return ""
	}
}

// 通过正则从字符串中提取匹配的多个值
func ExtractMultipleString(str string, regexpStr string) []string {
	re := regexp.MustCompile(regexpStr)
	return re.FindStringSubmatch(str)
}

// 通过正则统计出现次数
func CountMatches(str string, regexpStr string) int {
	regex := regexp.MustCompile(regexpStr)
	matches := regex.FindAllString(str, -1)
	return len(matches)
}
