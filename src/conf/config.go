package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type AgfunConf struct {
	SysDB  string `yaml:"sys_db"`
	AuthDB string `yaml:"auth_db"`
}

var conf *AgfunConf

func AgfunInst() *AgfunConf {
	if conf == nil {
		data, _ := ioutil.ReadFile("D:/myPro/go-service/src/conf/config.yml")
		fmt.Println(string(data))
		t := AgfunConf{}
		//把yaml形式的字符串解析成struct类型
		yaml.Unmarshal(data, &t)
		fmt.Println("初始数据", t)
		if t.SysDB == "" {
			fmt.Println("配置文件设置错误")
			return nil
		}
		conf = &t
	}
	return conf
}
