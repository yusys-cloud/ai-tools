// Author: yangzq80@gmail.com
// Date: 2023/8/16
package http

import (
	log "github.com/sirupsen/logrus"
	"github.com/yusys-cloud/ai-tools/base/flow/step/cmd"
	"github.com/yusys-cloud/ai-tools/base/search"
	"github.com/yusys-cloud/ai-tools/utils"
	"strings"
)

type Step struct {
	Url    string
	Header map[string]string
	Resp   *Resp
}

func (h *Step) Exec(log *log.Logger) map[string]interface{} {

	respJ := utils.Get(h.Url, h.Header)

	// 变量解析赋值处理
	if h.Resp.Cmd != nil {
		oldcmd := h.Resp.Cmd.Cmd
		v := search.ExtractAllString(h.Resp.Cmd.Cmd, `\$(\w+)`)
		for _, i := range respJ["items"].([]interface{}) {
			h.Resp.Cmd.Cmd = oldcmd
			im := i.(map[string]interface{})
			for _, s := range v {
				h.Resp.Cmd.Cmd = strings.ReplaceAll(h.Resp.Cmd.Cmd, "$"+s, im[s].(string))
			}
			h.Resp.Cmd.Exec(log)
		}
	}

	return respJ
}

type Resp struct {
	Cmd *cmd.Step
}
