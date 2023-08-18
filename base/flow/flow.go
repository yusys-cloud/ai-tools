// Author: yangzq80@gmail.com
// Date: 2023/6/14
package flow

import (
	"github.com/yusys-cloud/ai-tools/base/conf"
	"github.com/yusys-cloud/ai-tools/base/flow/step/cmd"
	"github.com/yusys-cloud/ai-tools/base/logs"
)

type Flow struct {
	Name  string
	Steps []*Step
	Cmds  []string //可选模式
}

func New(confName string) *Flow {
	flow := Flow{}
	conf.LoadJsonConfigFile(confName, &flow)
	return &flow
}

func (f *Flow) Run() {
	logger := logs.NewLoggerOutputFile(true, f.Name+".log")
	if f.Cmds != nil {
		for _, c := range f.Cmds {
			step := &cmd.Step{
				Cmd: c,
			}
			step.Exec(logger)
		}
	}
	for _, step := range f.Steps {
		step.exec(logger)
	}
	logger.Info("Flow end...")
}
