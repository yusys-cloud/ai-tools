// Author: yangzq80@gmail.com
// Date: 2021-02-23
//
package db

import (
	"fmt"
	"github.com/xujiajun/utils/filesystem"
	"os"
	"testing"
)

func TestA(t *testing.T) {

}

func TestCmd(t *testing.T) {
	//ExecCommand("vue init webpack demo")
	bptRootIdxDir := "aa/b"
	if ok := filesystem.PathIsExist(bptRootIdxDir); !ok {
		if err := os.MkdirAll(bptRootIdxDir, os.ModePerm); err != nil {
			fmt.Println(err)
			//return nil, err
		}
	}
}
