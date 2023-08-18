// Author: yangzq80@gmail.com
// Date: 2023/8/17
package text

import (
	log "github.com/sirupsen/logrus"
	"github.com/yusys-cloud/ai-tools/base/flow/step"
	"github.com/yusys-cloud/ai-tools/base/flow/step/http"
	"github.com/yusys-cloud/ai-tools/pkg/utils/convert"
	utils "github.com/yusys-cloud/ai-tools/pkg/utils/text"
	"strings"
)

type Step struct {
	Path      string
	Delimiter string
	Http      *http.Http
}

func (s *Step) Exec(log *log.Logger) {

	rawVar := s.Http.Payload
	vars := step.GetVariable(rawVar)

	utils.ScanTextLine(s.Path, func(line string, i int) bool {
		parts := strings.Split(line, s.Delimiter)

		s.Http.Payload = rawVar
		for _, v := range vars {
			s.Http.Payload = strings.ReplaceAll(s.Http.Payload, "$"+v, strings.TrimSpace(parts[convert.StrToInt(v)]))
		}
		//fmt.Println(s.Http.Payload)
		s.Http.Do()
		return true
	})
}