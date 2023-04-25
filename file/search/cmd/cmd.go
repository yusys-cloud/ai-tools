// Author: yangzq80@gmail.com
// Date: 2023/3/9
package main

import (
	"flag"
	"github.com/yusys-cloud/ai-tools/file/search"
)

func main() {
	path := flag.String("path", "conf.json", "--path=conf.json")
	flag.Parse()

	cnf := search.NewFileConf(*path)
	cnf.SearchFile.WalkContent(cnf.RootDir)
}
