// Author: yangzq80@gmail.com
// Date: 2023/7/21
package cmd

import (
	"os"
	"testing"
)

func TestH(t *testing.T) {
	os.Args = []string{"main", "-h"}
	Execute()
}

func TestJsonAPI(t *testing.T) {
	os.Args = []string{"main", "jsonapi"}
	Execute()
}

func TestWebSSH(t *testing.T) {
	os.Args = []string{"main", "webSSH"}
	Execute()
}

func TestOPSHelp(t *testing.T) {
	os.Args = []string{"main", "ops", "-h"}
	Execute()
}
func TestOPS(t *testing.T) {
	os.Args = []string{"main", "ops", "-k", "java"}
	Execute()
}
