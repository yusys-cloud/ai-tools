// Author: yangzq80@gmail.com
// Date: 2021-03-15
//
package server

import (
	"fmt"
	"log"
	"os/exec"
)

type Executor struct {
}

func ExecCommand(cmds string) {

	cmd := exec.Command("/bin/bash", "-c", cmds)

	out, err := cmd.Output()
	if err != nil {
		log.Fatal("error", err)
	}

	fmt.Println(string(out))
}
