// Author: yangzq80@gmail.com
// Date: 2021-02-23
//
package db

import (
	"testing"
)

func TestA(t *testing.T) {

	s := NewStorage("/tmp/ai-tools")
	s.Save("frontend", "vue-project-create", "vue工程创建")

	s.Save("ssh", "node1", "ssh root@n1")
	s.Save("ssh", "node2", "ssh root@n1")

	s.GetAll("frontend")

	s.GetAll("ssh")
}

func TestCmd(t *testing.T) {
	//ExecCommand("vue init webpack demo")
}
