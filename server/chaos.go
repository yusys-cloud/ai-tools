// Author: yangzq80@gmail.com
// Date: 2021-04-01
//
package server

import (
	"fmt"
	"github.com/Andrew-M-C/go.jsonvalue"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (s *Server) run(c *gin.Context) {

	kv := s.db.ReadOneRaw(c.Param("b"), c.Param("kid"))

	j, _ := jsonvalue.Unmarshal(kv)

	chaos, _ := j.GetArray("chaos")

	chaos.RangeArray(func(i int, v *jsonvalue.V) bool {
		node, _ := v.Get("node")
		ip, _ := node.GetString("ip")

		blades, _ := v.GetArray("blades")

		blades.RangeArray(func(i int, v *jsonvalue.V) bool {
			cmd, _ := v.GetString("cmd")

			//拼命令
			req := cmd
			logrus.Info(ip, "---", i, "---", cmd)

			if cmd == "thread-wait" {
				p, _ := v.GetArray("params")
				p.RangeArray(func(i int, v *jsonvalue.V) bool {
					t, _ := v.GetString("value")
					it, _ := strconv.Atoi(t)
					time.Sleep(time.Duration(it) * time.Second)
					return false
				})
				return true
			}

			bladeParams, _ := v.GetArray("params")

			//拼命令[参数]
			bladeParams.RangeArray(func(i int, v *jsonvalue.V) bool {

				id, _ := v.GetString("id")
				idv, _ := v.GetString("value")

				if idv != "" {

					req = fmt.Sprintf("%s %s %s", req, id, idv)

				}

				return true
			})

			req = url.PathEscape(req)
			//默认拼接:6666端口
			var freq string
			if !strings.Contains(ip, ":") {
				freq = fmt.Sprintf("http://%s:6666/chaosblade?cmd=%s", ip, req)
			} else {
				freq = fmt.Sprintf("http://%s/chaosblade?cmd=%s", ip, req)
			}
			//执行blade request
			logrus.Println("send...", freq)
			rs := jsonvalue.NewObject()
			rs.SetString(freq).At("reqUrl")
			//rs.SetString(getUrl(freq)).At("respBody")

			v.Set(rs).At("runResult")
			//保存执行结果到json
			logrus.Println(chaos)

			return true
		})

		return true
	})
	s.db.UpdateMarshalValue(c.Param("b"), c.Param("kid"), j.MustMarshal())
}

func getUrl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error(err.Error())
		return "http.get i/o timeout"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}

	return string(body)
}
