// Author: yangzq80@gmail.com
// Date: 2023/3/8
package search

import (
	"testing"
)

func TestShowStruct(t *testing.T) {
	search := NewSearch("conf.json")
	search.Rule.WalkContent(search.RootDir)
}
