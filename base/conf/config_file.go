// Author: yangzq80@gmail.com
// Date: 2023/3/27
// Load the specified JSON file into a struct.
// eg:
// func NewFileConf(name string) *Conf {
//	cnf := &Conf{}
//	file.LoadJsonConfigFile(name, cnf)
//	return cnf
// }

package conf

import (
	"encoding/json"
	"log"
	"os"
)

// 从json文件中加载配置到struct中，如 :
// search := &Search{} ;
// conf.LoadJsonConfigFile("conf.json", search)
func LoadJsonConfigFile(jsonFilename string, confStruct *interface{}) {
	data, err := os.ReadFile(jsonFilename)

	if err != nil {
		log.Panicf("read config file [%s] failure:%+v", jsonFilename, err)
	}
	err = json.Unmarshal(data, confStruct)
	if err != nil {
		log.Panicf("parse config file [%s] failure. error:%+v", jsonFilename, err)
	}
	log.Printf("Init config...%v\n---Init config end---\n", string(data))

}
