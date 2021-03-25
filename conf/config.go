// Author: yangzq80@gmail.com
// Date: 2021-03-25
//
package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Conf struct {
	Port string `json:"port,omitempty"`
	Path string `json:"path,omitempty"`
	Mode string
}

func ReadConfig() *Conf {
	file := "config.json"

	data, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("read config file <%s> failure. err:%+v", file, err)
	}

	cnf := &Conf{}
	err = json.Unmarshal(data, cnf)
	if err != nil {
		log.Fatalf("parse config file <%s> failure. error:%+v", file, err)
	}

	return cnf
}
