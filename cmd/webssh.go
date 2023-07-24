// Author: yangzq80@gmail.com
// Date: 2023/7/18
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yusys-cloud/ai-tools/conf"
	"github.com/yusys-cloud/ai-tools/server"
)

var sshCmd = &cobra.Command{
	Use:     "webSSH",
	Short:   "可视化SSH API 服务",
	Long:    `可视化操作常用ssh命令`,
	Example: `./ai-tools webSSH`,
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cf := &conf.Conf{port, dataPath, "dev"}

		s := server.NewServer(cf)

		s.Start()
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
	sshCmd.Flags().StringVarP(&dataPath, "data", "d", "./data", "JSON数据存储目录")
	sshCmd.Flags().StringVarP(&port, "port", "p", "9999", "HTTP端口")
}
