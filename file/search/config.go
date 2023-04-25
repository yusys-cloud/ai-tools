// Author: yangzq80@gmail.com
// Date: 2023/3/9
package search

import "github.com/yusys-cloud/ai-tools/file/conf"

type Conf struct {
	RootDir    string
	SearchFile *SearchFile
}
type SearchFile struct {
	Suffix []string
	//排除含有指定路径的文件
	Exclude []string
	Content *Content
	Output  *Output
}
type Content struct {
	Include []string
	Exclude []string //支持正则
	//Regexp  []string
	Replace string //替换为新的字符串
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

func NewFileConf(name string) *Conf {
	cnf := &Conf{}
	conf.LoadJsonConfigFile(name, cnf)
	return cnf
}
