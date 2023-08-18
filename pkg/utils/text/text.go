// Author: yangzq80@gmail.com
// Date: 2023/7/21
package text

import (
	"bufio"
	"log"
	"os"
)

// 返回false则停止扫描
type TextLineFunc func(line string, i int) bool

func ScanTextLine(filepath string, lineFunc TextLineFunc) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println("ScanTextLine Error:", err)
		return err
	}

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if !lineFunc(line, i) {
			break
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Println("ScanTextLine Error:", err)
		return err
	}
	return nil
}
