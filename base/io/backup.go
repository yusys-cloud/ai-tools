// Author: yangzq80@gmail.com
// Date: 2023/4/25
package io

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func BackupFolder(sourceDir string) {
	i := 0
	backupDir := fmt.Sprintf("%s-backup-%s", sourceDir, time.Now().Format("2006-01-02-150405"))

	err := filepath.Walk(sourceDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 如果是目录，则创建备份目录
		if info.IsDir() {
			targetDir := filepath.Join(backupDir, path[len(sourceDir):])
			return os.MkdirAll(targetDir, 0755)
		}

		// 如果是文件，则拷贝到备份目录中
		if !info.Mode().IsRegular() {
			return nil
		}

		sourceFile, err := os.Open(path)
		if err != nil {
			return err
		}

		i++
		targetFile := filepath.Join(backupDir, path[len(sourceDir):])
		return copyFile(sourceFile, targetFile, info)
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%v files backup completed successfully.", i)
}

func copyFile(sourceFile *os.File, targetPath string, sourceFileinfo fs.FileInfo) error {
	// 打开目标文件，如果不存在则创建
	targetFile, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	// 拷贝数据
	_, err = io.Copy(targetFile, sourceFile)
	if err != nil {
		return err
	}

	// 拷贝文件权限和时间戳
	targetFile.SetWriteDeadline(sourceFileinfo.ModTime())

	return os.Chmod(targetPath, sourceFileinfo.Mode())
}
