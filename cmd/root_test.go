// Author: yangzq80@gmail.com
// Date: 2023/7/21
package cmd

import (
	"os"
	"testing"
)

func TestJsonAPI(t *testing.T) {
	os.Args = []string{"main", "jsonapi"}
	Execute()
}

func TestWebSSH(t *testing.T) {
	os.Args = []string{"main", "webSSH"}
	Execute()
}
