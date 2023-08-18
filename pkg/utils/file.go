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
