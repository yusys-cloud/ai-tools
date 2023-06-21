// Author: yangzq80@gmail.com
// Date: 2023/3/8
package search

import (
	"fmt"
	"testing"
)

var search = NewSearch("_test_conf.json")

func TestShowStruct(t *testing.T) {
	search.Rule.WalkContent(search.RootDir)
}

func TestExtContentFunc(t *testing.T) {
	search.Rule.WalkContentWithFunc(search.RootDir, nil, nil, func(f ContentExtMatchedFunc, path string, content string) {
		fmt.Println(f.FuncName, f.Include)
	}, nil)
}

func TestRename(t *testing.T) {
	NewSearch("_test_rename_conf.json").Start()
}
