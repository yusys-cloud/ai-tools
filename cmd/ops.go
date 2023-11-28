// Author: yangzq80@gmail.com
// Date: 2023/8/18
// 减少Linux 命令输入
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func init() {

	var killName string
	var killConfirm string

	var opsCmd = &cobra.Command{
		Use:   "ops",
		Short: "运维相关命令",
		Long:  `流程执行`,
		Run: func(cmd *cobra.Command, args []string) {
			if killConfirm == "n" {
				killStr := fmt.Sprintf("kill -9 $(ps -ef | grep %v | grep -v grep | awk '{ print $2 }')", killName)
				rs, _ := exec.Command("bash", "-c", killStr).CombinedOutput()
				fmt.Println(string(rs))
				return
			}
			output, err := exec.Command("bash", "-c", "ps aux | grep "+killName).CombinedOutput()
			if err != nil {
				fmt.Println("Error executing command:", err)
				return
			}

			processList := strings.Split(string(output), "\n")
			processList = processList[3:]
			if len(processList) == 1 && len(processList[0]) == 0 {
				fmt.Println("Not found :" + killName)
				return
			}

			// 显示带有序号的进程列表
			i := 0
			for _, process := range processList {
				if len(process) > 0 {
					i++
					fmt.Printf("%d. %s\n", i, process)
				}
			}

			// 获取用户输入的进程序号
			input := userInput("Enter the process number to kill:")
			num, err := strconv.Atoi(input)

			// 获取对应序号的进程PID
			pid := getPIDFromProcessString(processList[num-1])

			fmt.Sprintf("Confirm kill PID: %v  y/n \n", pid)

			// 执行 kill 命令
			killCmd := exec.Command("kill", pid)
			killOutput, killErr := killCmd.CombinedOutput()
			if killErr != nil {
				fmt.Println("Error killing process:", killErr.Error())
				return
			}
			fmt.Println(string(killOutput))
		},
	}

	opsCmd.Flags().StringVarP(&killName, "kill", "k", "", "-k gs-spring.jar //进程名称")
	opsCmd.Flags().StringVarP(&killConfirm, "confirm", "c", "n", "-c y/n //是否进行操作确认")

	rootCmd.AddCommand(opsCmd)
}

// 用户从命令行输入字符
func userInput(tips string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(tips)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// 从进程字符串中获取PID
func getPIDFromProcessString(process string) string {
	fields := strings.Fields(process)
	if len(fields) > 1 {
		return fields[1]
	}
	return ""
}
