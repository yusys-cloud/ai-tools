// Author: yangzq80@gmail.com
// Date: 2023/7/21
package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 返回false则停止扫描
type TextLineFunc func(line string, i int) bool

// 对filepath文本进行逐行读取，自定义TextLineFunc方法处理每一行数据
func ScanTextLine(filepath string, lineFunc TextLineFunc) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println("ScanTextLine-openFile Error:", err)
		return err
	}

	r := bufio.NewReader(file)
	i := 0
	for {
		// 读取一行数据
		line, isPrefix, err := r.ReadLine()

		if err != nil {
			//fmt.Println("Error reading file:", err)
			break
		}

		i++
		if !lineFunc(string(line), i) {
			break
		}

		if isPrefix {
			fmt.Println("Line is too long and continues on the next line.", string(line))
		}
	}

	//scanner := bufio.NewScanner(file)
	//
	//buf := make([]byte, 0, 64*1024)
	//scanner.Buffer(buf, 1024*1024)
	//
	//i := 0
	//for scanner.Scan() {
	//	line := scanner.Text()
	//	if !lineFunc(line, i) {
	//		break
	//	}
	//	i++
	//}
	//
	//if err := scanner.Err(); err != nil {
	//	log.Println("ScanTextLine Error:", err, filepath)
	//	return err
	//}
	return nil
}
