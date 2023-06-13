// Author: yangzq80@gmail.com
// Date: 2023/6/13
package logs

import (
	log "github.com/sirupsen/logrus"
	"github.com/yusys-cloud/ai-tools/base/time"
	"io"
	"os"
)

func NewLogger(debug bool) *log.Logger {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{TimestampFormat: time.DateTimeFormat})
	if debug {
		logger.SetLevel(log.DebugLevel)
	}
	return logger
}
func NewLoggerOutputFile(debug bool, name string) *log.Logger {
	logger := NewLogger(debug)
	// 设置日志输出为文件
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		// 日志输出同时定向到控制台和文件
		logger.SetOutput(io.MultiWriter(os.Stdout, file))
	} else {
		logger.Error("Failed to open log file, using default stderr", err.Error())
	}
	return logger
}
