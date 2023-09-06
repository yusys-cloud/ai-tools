// Author: yangzq80@gmail.com
// Date: 2023/8/18
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yusys-cloud/ai-tools/base/flow"
)

var flowCmd = &cobra.Command{
	Use:   "flow",
	Short: "自动化流程执行",
	Long:  `流程执行`,
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		flow.New(path).Run()
	},
}
var path string

func init() {
	rootCmd.AddCommand(flowCmd)
	flowCmd.Flags().StringVarP(&path, "config", "c", "config.json", "启动配置文件")
}
