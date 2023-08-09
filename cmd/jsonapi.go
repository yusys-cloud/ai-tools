// Author: yangzq80@gmail.com
// Date: 2023/7/18
package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/yusys-cloud/go-jsonstore-rest/rest"
)

var apisCmd = &cobra.Command{
	Use:   "jsonapi",
	Short: "开箱即用的 JSON 文件存储 REST API 服务",
	Long: `这是一款开箱即用的 JSON 文件存储 REST API 服务。它提供了一个简单的、无需额外数据库的方式来创建、读取、更新和删除 JSON 数据。
你可以通过发送 HTTP 请求来与这个服务交互，它会将你提供的 JSON 数据存储在文件中，并根据请求返回相应的结果。
这个服务特别适用于小规模的应用或原型开发，让你能够快速搭建一个支持数据存储的 REST API 环境`,
	Example: `新增:	curl localhost:9999/kv/meta/node -X POST -d '{"ip": "192.168.x.x","name":"redis-n1"}' -H "Content-Type: application/json"
修改:	curl localhost:9999/kv/meta/node/node:1429991523109310464 -X PUT -d '{"ip": "192.168.49.69","name":"redis-n2"}' -H "Content-Type: application/json"
查看:	curl localhost:9999/kv/meta/node
删除:	curl localhost:9999/kv/meta/node/node:1429991523109310464 -X DELETE
`,
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		//REST-APIs-json
		s := rest.NewJsonStoreRest(dataPath)
		s.ConfigHandles(r)
		s.DisableCors = true
		r.Run(":" + port)
	},
}

func init() {
	rootCmd.AddCommand(apisCmd)
	apisCmd.Flags().StringVarP(&dataPath, "data", "d", "./data", "JSON数据存储目录")
	apisCmd.Flags().StringVarP(&port, "port", "p", "9999", "HTTP端口")
}
