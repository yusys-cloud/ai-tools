// Author: yangzq80@gmail.com
// Date: 2023/8/17
package utils

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestScanTextLine(t *testing.T) {
	filepath := "/Users/zqy/test/tmp/test.ts"

	var history []string

	ScanTextLine(filepath, func(line string, i int) bool {

		history = append(history, line)
		fmt.Println(i, line)
		return true
	})
}

func TestScanTextLine_longLine(t *testing.T) {
	tmpFile := "long_text_file.txt"
	// 生成一个很长的行
	longLine := "1" + strings.Repeat("a", 10000) + "\n" + strings.Repeat("b", 10000)

	// 将生成的行写入文件
	file, err := os.Create(tmpFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	_, err = file.WriteString(longLine)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	ScanTextLine(tmpFile, func(line string, i int) bool {
		fmt.Println(len(line), line)
		return true
	})
	os.RemoveAll(tmpFile)
}
