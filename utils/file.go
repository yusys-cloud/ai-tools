// Author: yangzq80@gmail.com
// Date: 2023/7/21
package utils

import "os"

// Exists determine whether the file exists
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// FileExists returns true if file exists.
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

// OpenOrCreateFile opens file or creates it if it doesn't exist.
func OpenOrCreateFile(file string) (*os.File, error) {
	if FileExists(file) {
		outfile, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0755)
		if err != nil {
			return nil, err
		}
		return outfile, nil
	} else {
		outfile, err := os.Create(file)
		if err != nil {
			return nil, err
		}
		return outfile, nil
	}
}
