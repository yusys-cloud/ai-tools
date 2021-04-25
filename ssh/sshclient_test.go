// Author: yangzq80@gmail.com
// Date: 2021-04-19
//
package ssh

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"testing"
)

var (
	addr   string
	user   string
	passwd string
	prikey string
)

func TestDailWithPasswd(t *testing.T) {
	client, err := DialWithPasswd(addr, user, passwd)
	if err != nil {
		t.Fatal("DialWithPasswd err: ", err)
	}
	if err := client.Close(); err != nil {
		t.Fatal("client.Close err: ", err)
	}
}

func TestDailWithKey(t *testing.T) {
	client, err := DialWithKey(addr, user, prikey)
	if err != nil {
		t.Fatal("DialWithPasswd err: ", err)
	}
	if err := client.Close(); err != nil {
		t.Fatal("client.Close err: ", err)
	}
}

func TestCmdRun(t *testing.T) {
	client, err := DialWithKey(addr, user, prikey)
	if err != nil {
		t.Fatal("DialWithPasswd err: ", err)
	}
	defer client.Close()

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	err = client.Cmd("echo stdout").Cmd(">&2 echo stderr").SetStdio(&stdout, &stderr).Run()
	if err != nil {
		t.Fatal("Run command err: ", err)
	}

	fmt.Println(stdout.String())

	if stdout.String() != "stdout\n" {
		t.Fatal("Command output mismatching on stdout")
	}
	if stderr.String() != "stderr\n" {
		t.Fatal("Command output mismatching on stderr")
	}
}

func TestShell(t *testing.T) {
	client, err := DialWithKey(addr, user, prikey)
	if err != nil {
		t.Fatal("DialWithPasswd err: ", err)
	}
	defer client.Close()

	script := bytes.NewBufferString("echo stdout\n  >&2 echo stderr")
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)
	err = client.Shell().SetStdio(script, &stdout, &stderr).Start()
	if err != nil {
		t.Fatal("Start shell faield: ", err)
	}

	if stdout.String() != "stdout" == false {
		t.Fatal("Command output mismatching on stdout")
	}
	if stderr.String() != "stderr" == false {
		t.Fatal("Command output mismatching on stderr")
	}
}

func TestMain(m *testing.M) {
	flag.StringVar(&addr, "addr", "food:22", "The host of ssh")
	flag.StringVar(&user, "user", "root", "The user of login")
	flag.StringVar(&passwd, "passwd", "yourpasswd", "The passwd of user")
	flag.StringVar(&prikey, "privatekey", "/Users/zqy/.ssh/id_rsa", "The privatekey of user")

	flag.Parse()
	os.Exit(m.Run())
}
