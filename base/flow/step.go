// Author: yangzq80@gmail.com
// Date: 2023/6/14
package flow

import (
	log "github.com/sirupsen/logrus"
	"github.com/yusys-cloud/ai-tools/base/flow/step/cmd"
	"github.com/yusys-cloud/ai-tools/base/flow/step/http"
	"github.com/yusys-cloud/ai-tools/base/flow/step/text"
)

type Step struct {
	Cmd  *cmd.Step
	Http *http.Step
	Text *text.Step
}

func (s *Step) exec(log *log.Logger) error {
	if s.Http != nil {
		s.Http.Exec(log)
	}
	if s.Cmd != nil {
		s.Cmd.Exec(log)
	}
	if s.Text != nil {
		s.Text.Exec(log)
	}
	return nil
}
