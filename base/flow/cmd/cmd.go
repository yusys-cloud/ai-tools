// Author: yangzq80@gmail.com
// Date: 2023/6/15
package main

import (
	"flag"
	"github.com/yusys-cloud/ai-tools/base/flow"
)

func main() {
	path := flag.String("path", "conf.json", "--path=conf.json")
	flag.Parse()
	flow.New(*path).Run()
}
