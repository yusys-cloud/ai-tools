// Author: yangzq80@gmail.com
// Date: 2023/8/17
package step

import "github.com/yusys-cloud/ai-tools/base/search"

// 抽取字符串中的$变量名
func GetVariable(str string) []string {
	return search.ExtractAllString(str, `\$(\w+)`)
}
