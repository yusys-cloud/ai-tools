// Author: yangzq80@gmail.com
// Date: 2023/3/9
package search

import "github.com/yusys-cloud/ai-tools/base/conf"

type Search struct {
	RootDir string
	Rule    *Rule
}
type Rule struct {
	Suffix []string
	//排除含有指定路径的文件
	Exclude []string
	Content *Content
	// 自定义对全部内容处理
	ContentExtFunc []*ContentExtMatchedFunc
	// 自定义对匹配到的内容处理
	ContentExtMatchedFunc []*ContentExtMatchedFunc
}
type ContentExtMatchedFunc struct {
	FuncName string
	Include  []string
}
type Content struct {
	Include []string
	Exclude []string //支持正则
	Replace string   //替换为新的字符串
	Output  *Output
}
type Output struct {
	FileEnable    bool
	ContentEnable bool
	ContentShow   *ContentShow
}
type ContentShow struct {
	All           bool
	AdjacentLines int
}

func NewSearch(configFileName string) *Search {
	cnf := &Search{}
	conf.LoadJsonConfigFile(configFileName, cnf)
	return cnf
}
