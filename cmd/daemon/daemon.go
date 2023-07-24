// Author: yangzq80@gmail.com
// Date: 2023/7/21
package daemon

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

// StartCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Silent start alist server with `--force-bin-dir`",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		start("server")
	},
}

func start(dcmd string) {
	initDaemon()
	if pid != -1 {
		_, err := os.FindProcess(pid)
		if err == nil {
			log.Info("alist already started, pid ", pid)
			return
		}
	}
	args := os.Args
	args[1] = dcmd
	args = append(args, "--force-bin-dir")
	cmd := &exec.Cmd{
		Path: args[0],
		Args: args,
		Env:  os.Environ(),
	}
	stdout, err := os.OpenFile(filepath.Join(filepath.Dir(pidFile), "start.log"), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(os.Getpid(), ": failed to open start log file:", err)
	}
	cmd.Stderr = stdout
	cmd.Stdout = stdout
	err = cmd.Start()
	if err != nil {
		log.Fatal("failed to start children process: ", err)
	}
	log.Infof("success start pid: %d", cmd.Process.Pid)
	err = os.WriteFile(pidFile, []byte(strconv.Itoa(cmd.Process.Pid)), 0666)
	if err != nil {
		log.Warn("failed to record pid, you may not be able to stop the program with `./alist stop`")
	}
}

func init() {
	//rootCmd.AddCommand(StartCmd)
}
