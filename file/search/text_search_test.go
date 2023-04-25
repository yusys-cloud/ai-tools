// Author: yangzq80@gmail.com
// Date: 2023/3/8
package search

import (
	"testing"
)

func TestShowStruct(t *testing.T) {
	as := &Search{NewFileConf("conf-photo.json")}
	as.Conf.SearchFile.WalkContent(as.Conf.RootDir)
}
