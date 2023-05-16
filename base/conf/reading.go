// Author: yangzq80@gmail.com
// Date: 2023/5/6
package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

// 从配置文件读取值 如:yaml中server.port为key取值
// 支持 JSON, TOML, YAML, HCL, envfile and Java properties config files
func ReadFromConfFile(fileName string, key string) (interface{}, *viper.Viper) {
	v := viper.New()
	v.SetConfigFile(fileName)
	err := v.ReadInConfig()
	if err != nil {
		fmt.Sprintf("error reading config: %s", err)
		return nil, nil
	}
	return v.Get(key), v
}
