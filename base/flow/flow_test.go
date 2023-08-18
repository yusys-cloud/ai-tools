// Author: yangzq80@gmail.com
// Date: 2023/6/14
package flow

import (
	"fmt"
	"github.com/yusys-cloud/ai-tools/base/search"
	"log"
	"os/exec"
	"testing"
)

func TestStepHttp(t *testing.T) {
	v := search.ExtractAllString("git clone $url $branch", `\$(\w+)`)
	fmt.Println(v)
	New("step/http/_test_conf.json").Run()
}

func TestStepTextToHttp(t *testing.T) {
	New("step/text/_test_conf.json").Run()
}

func TestStepCmd(t *testing.T) {
	New("step/cmd/_test_conf.json").Run()
}

func TestExec(t *testing.T) {
	cmd := exec.Command("bash", "-c", "cd /Users/zqy/test/tmp && pwd && ls&&python3 -V")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(out))
}
