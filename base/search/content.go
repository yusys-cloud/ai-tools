// Author: yangzq80@gmail.com
// Date: 2023/3/8
package search

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func defaultContentHandle(sr *Rule, matchedFunc HandleMatchedFunc, path string, data []byte, content string, err error, info os.FileInfo, ct int, rootDir string) {
	//print-file-path
	if sr.Content.Output.FileEnable {
		fmt.Printf("------TargetFile[%v]------%v \n", ct, strings.ReplaceAll(path, rootDir, "")) // 显示文件名
	}
	if matchedFunc != nil {
		matchedFunc(path, data)
	}
	//print-content
	if sr.Content.Output.ContentEnable {
		if sr.Content.Output.ContentShow.All {
			fmt.Println(content)
		} else {
			showAdjacent(path, sr.Content.Output.ContentShow.AdjacentLines, sr.Content.Include)
		}
	}
	// replace
	if sr.Content.Replace != "" {
		modifiedContent := ""
		// replace all occurrences of "abc" with "aabbcc"
		for _, s := range sr.Content.Include {
			modifiedContent = strings.ReplaceAll(content, s, sr.Content.Replace)
		}

		// write the modified string back to the file
		err = os.WriteFile(path, []byte(modifiedContent), info.Mode())
		if err != nil {
			fmt.Println(err.Error())
			//return nil
		}
	}
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
