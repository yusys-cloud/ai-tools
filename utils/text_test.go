// Author: yangzq80@gmail.com
// Date: 2023/8/17
package utils

import (
	"fmt"
	"testing"
)

type Method struct {
	Declaration string
}

func TestScanTextLine(t *testing.T) {
	filepath := "/Users/zqy/test/tmp/test.ts"

	var history []string

	ScanTextLine(filepath, func(line string, i int) bool {

		history = append(history, line)
		fmt.Println(i, line)
		return true
	})
}
