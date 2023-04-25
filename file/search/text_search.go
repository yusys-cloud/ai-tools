// Author: yangzq80@gmail.com
// Date: 2023/3/8
package search

import (
	"bufio"
	"fmt"
	"github.com/yusys-cloud/ai-tools/file/io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Search struct {
	Conf *Conf
}

type PathFunc func(path string) bool
type HandleMatchedFunc func(path string, content []byte)

func (sf *SearchFile) WalkContent(rootDir string) error {
	if sf.Content.Replace != "" {
		io.MakeBackup(rootDir)
	}
	return sf.WalkContentWithFunc(rootDir, nil, nil)
}

func (tf *SearchFile) WalkContentWithFunc(rootDir string, pathFunc PathFunc, matchedFunc HandleMatchedFunc) error {
	cAll := 0
	ct := 0
	//tf := s.Conf.SearchFile
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		cAll++
		if err != nil {
			return err
		}
		if pathFunc != nil {
			if !pathFunc(path) {
				return nil
			}
		}
		if info.IsDir() {
			return nil
		}
		// 排除含有指定路径的文件
		if ContainsStr(path, tf.Exclude) {
			return nil
		}
		// 帅选指定类型后缀的文件
		if len(tf.Suffix) > 0 {
			ext := filepath.Ext(path)
			if !ContainsStr(ext, tf.Suffix) {
				return nil
			}
		}
		// 读取并输出文件内容
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		content := string(data)

		// 不包含排除字符串
		if ContainsStr(content, tf.Content.Exclude) {
			return nil
		}
		// 包含指定字符串
		if ContainsStr(content, tf.Content.Include) {
			ct++
			//print-file-path
			if tf.Output.FileEnable {
				fmt.Printf("------TargetFile[%v]------%v \n", ct, strings.ReplaceAll(path, rootDir, "")) // 显示文件名
			}
			if matchedFunc != nil {
				matchedFunc(path, data)
			}
			//print-content
			if tf.Output.ContentEnable {
				if tf.Output.ContentShow.All {
					fmt.Println(content)
				} else {
					showAdjacent(path, tf.Output.ContentShow.AdjacentLines, tf.Content.Include)
				}
			}
			// replace
			if tf.Content.Replace != "" {
				modifiedContent := ""
				// replace all occurrences of "abc" with "aabbcc"
				for _, s := range tf.Content.Include {
					modifiedContent = strings.ReplaceAll(content, s, tf.Content.Replace)
				}

				// write the modified string back to the file
				err = os.WriteFile(path, []byte(modifiedContent), info.Mode())
				if err != nil {
					return nil
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("------TotalFile:%v targetFile:%v------ \n", cAll, ct)
	return nil
}

func ContainsStr(content string, strs []string) bool {
	for _, e := range strs {
		if regexp.MustCompile(e).MatchString(content) {
			//if strings.Contains(content, e) {
			return true
		}
	}
	return false
}

// 显示临近行，0 则显示本身
func showAdjacent(path string, adjacentLines int, include []string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if ContainsStr(line, include) {
			fmt.Println("```")
			if adjacentLines == 0 {
				fmt.Printf("%d %s\n", lineNum, line)
			} else {
				showAdjacentLines(path, lineNum, adjacentLines)
			}
			fmt.Println("```")
		}
	}
}

func showAdjacentLines(path string, lineNum int, adjacentLines int) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	startLine := lineNum - adjacentLines
	if startLine < 1 {
		startLine = 1
	}

	endLine := lineNum + adjacentLines

	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan() && i <= endLine; i++ {
		if i >= startLine {
			fmt.Printf("%d %s\n", i, scanner.Text())
		}
	}
}
