// Author: yangzq80@gmail.com
// Date: 2023/3/8
package search

import (
	"fmt"
	"github.com/yusys-cloud/ai-tools/base/io"
	"os"
	"path/filepath"
	"time"
)

// 自定义规则判断文件是否满足搜索条件，false则不满足，默认true
type PathFunc func(path string) bool

// 自定义匹配到的内容处理
type HandleMatchedFunc func(path string, content string)

// 扩展自定义func对全部内容进行处理
type ExtContentFunc func(contentExtFunc ContentExtMatchedFunc, path string, content string)

// 扩展自定义func对匹配到的内容进行处理
type ExtContentMatchedFunc func(contentExtMatchedFunc ContentExtMatchedFunc, path string, content string)

func (sf *Rule) WalkContent(rootDir string) error {
	return sf.WalkContentWithFunc(rootDir, nil, nil, nil, nil)
}

func (sr *Rule) WalkContentWithFunc(rootDir string, pathFunc PathFunc, matchedFunc HandleMatchedFunc, extContentFunc ExtContentFunc, extContentMatchedFunc ExtContentMatchedFunc) error {
	cAll := 0
	ct := 0
	start := time.Now()
	// 有字符串替换操作则先备份原目录 rootDir
	if sr.Content != nil && sr.Content.Replace != "" {
		io.BackupFolder(rootDir)
	}
	//tf := s.Conf.SearchFile
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		// s1.路径规则
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
		if ContainsStr(path, sr.Exclude) {
			return nil
		}
		// 帅选指定类型后缀的文件
		if len(sr.Suffix) > 0 {
			ext := filepath.Ext(path)
			if !ContainsStr(ext, sr.Suffix) {
				return nil
			}
		}
		// 读取并输出文件内容
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		content := string(data)

		// s2.根据配置规则进行搜索并显示
		if sr.Content != nil {
			// 不包含排除字符串
			if ContainsStr(content, sr.Content.Exclude) {
				return nil
			}
			// 包含指定字符串
			if sr.Content.Include == nil || len(sr.Content.Include) == 0 || ContainsStr(content, sr.Content.Include) {
				ct++
				//默认内容处理
				defaultContentHandle(sr, matchedFunc, path, content, err, info, ct, rootDir)
			}
		}

		// s3.扩展自定义内容处理
		if extContentFunc != nil {
			for _, f := range sr.ContentExtFunc {
				extContentFunc(*f, path, content)
			}
		}
		// s4.扩展自定义对匹配到内容处理
		if extContentMatchedFunc != nil {
			for _, f := range sr.ContentExtMatchedFunc {
				if ContainsStr(content, f.Include) {
					extContentMatchedFunc(*f, path, content)
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Search totalFile:%v targetFile:%v cost:%v \n", cAll, ct, time.Now().Sub(start))
	return nil
}
