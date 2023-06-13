// Author: yangzq80@gmail.com
// Date: 2023/6/13
package logs

import "testing"

func TestNewLogger(t *testing.T) {
	NewLogger(false).Debug("debug")
}

func TestNewLoggerOutputFile(t *testing.T) {
	NewLoggerOutputFile(true, "test.log").Debug("test debug")
}
