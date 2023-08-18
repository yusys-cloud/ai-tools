// Author: yangzq80@gmail.com
// Date: 2023/8/16
package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/yusys-cloud/ai-tools/base/search"
	"os/exec"
)

type Step struct {
	Cmd   string
	Retry *Retry
}

type Retry struct {
	Condition []string
	Times     int
}

func (s *Step) Exec(*log.Logger) {
	i := 0
	for {
		log.Infof("Exec[%v] start...", s.Cmd)
		cmd := exec.Command("/bin/bash", "-c", s.Cmd)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Infof("Exec cmd error:[%v] [%v]", err.Error(), s.Cmd)
		}

		if s.Retry != nil {
			log.Infof("Exec end output:[isRetry:%v %v]", search.ContainsStr(string(output), s.Retry.Condition), string(output))
			if output != nil && len(output) > 0 && i < s.Retry.Times {
				// 判断错误信息中是否包含特定字符串
				if search.ContainsStr(string(output), s.Retry.Condition) {
					i++
					log.Infof("Encountered an error. Retrying......【%v times】......", i)
					continue // 继续循环重新执行
				}
			}
		}
		// 执行成功，跳出循环
		break
	}
}
