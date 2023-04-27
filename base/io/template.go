// Author: yangzq80@gmail.com
// Date: 2021/8/26
package io

import (
	"os"
	"path/filepath"
	"text/template"
)

func GenerateTemplateFile(targetFileName string, data any, templateFileNames ...string) (*os.File, error) {
	tmpl, err := template.ParseFiles(templateFileNames...)
	if err != nil {
		return nil, err
	}
	err = CreateParentDir(targetFileName)
	if err != nil {
		return nil, err
	}
	file, err := os.Create(targetFileName)
	if err != nil {
		return nil, err
	}
	err = tmpl.Execute(file, data)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func CreateParentDir(filename string) error {
	parentDir := filepath.Dir(filename)
	if _, err := os.Stat(parentDir); os.IsNotExist(err) {
		err = os.MkdirAll(parentDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
