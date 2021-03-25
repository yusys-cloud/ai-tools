// Author: yangzq80@gmail.com
// Date: 2021-02-02
//
package main

import (
	"github.com/yusys-cloud/ai-tools/conf"
	"github.com/yusys-cloud/ai-tools/server"
)

func main() {
	cf := conf.ReadConfig()

	s := server.NewServer(cf)

	s.Start()
}
