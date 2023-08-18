// Author: yangzq80@gmail.com
// Date: 2021/3/28
package io

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// 将文件@filename上传到指定目录
func UploadFile(uploadPath, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening file:", err)
		return err
	}

	filePath := filepath.Join(uploadPath, filepath.Base(filename))
	os.MkdirAll(uploadPath, 0750)
	out, err := os.Create(filePath)
	if err != nil {
		log.Println("Error creating file:", err)
		return err
	}

	_, err = io.Copy(out, file)
	if err != nil {
		log.Println("Error copying file:", err)
		return err
	}

	log.Println("File uploaded successfully to:", filePath)

	return nil
}
