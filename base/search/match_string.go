// Author: yangzq80@gmail.com
// Date: 2021/10/6
package search

import (
	"regexp"
)

func ContainsStr(content string, strs []string) bool {
	for _, e := range strs {
		if regexp.MustCompile(e).MatchString(content) {
			//if strings.Contains(content, e) {
			return true
		}
	}
	return false
}

// 通过正则从字符串中提取匹配的值
func ExtractString(str string, regexpRule string) string {
	re := regexp.MustCompile(regexpRule)
	matches := re.FindStringSubmatch(str)

	if len(matches) > 1 {
		return matches[1]
	} else {
		return ""
	}
}
