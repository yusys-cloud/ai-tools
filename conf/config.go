// Author: yangzq80@gmail.com
// Date: 2021-03-25
package conf

import (
	"encoding/json"
	"log"
	"os"
)

type Conf struct {
	Port string `json:"port,omitempty"`
	Path string `json:"path,omitempty"`
	Mode string
}

func ReadConfig() *Conf {
	file := "config.json"

	data, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("read config file <%s> failure. err:%+v", file, err)
	}

	cnf := &Conf{}
	err = json.Unmarshal(data, cnf)
	if err != nil {
		log.Fatalf("parse config file <%s> failure. error:%+v", file, err)
	}

	b, _ := json.MarshalIndent(cnf, "", "   ")
	//b,_ := json.Marshal(conf)

	log.Println("Config:\n", string(b))

	return cnf
}
