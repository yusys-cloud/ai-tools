// Author: yangzq80@gmail.com
// Date: 2023/3/9
package search

import "github.com/yusys-cloud/ai-tools/base/conf"

type Search struct {
	RootDir    string
	SearchRule *SearchRule
}
type SearchRule struct {
	Suffix []string
	//排除含有指定路径的文件
	Exclude []string
	Content *Content
	Output  *Output
}
type Content struct {
	Include []string
	Exclude []string //支持正则
	Replace string   //替换为新的字符串
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
