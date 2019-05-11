package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type AgfunConf struct {
	SysDB     string   `yaml:"sys_db"`
	AuthDB    string   `yaml:"auth_db"`
	Etcd      []string `yaml:"etcd"`
	VideoHost string   `yaml: "video_host"`
}

var conf *AgfunConf

func AgfunInst() *AgfunConf {
	if conf == nil {
		data, _ := ioutil.ReadFile("C:/Users/fengc/Desktop/workspace/go-service/src/conf/configDebug.yml")
		fmt.Println(string(data))
		t := AgfunConf{}
		//把yaml形式的字符串解析成struct类型
		e := yaml.Unmarshal(data, &t)
		if e != nil {
			log.Fatal(e)
		}
		fmt.Println("初始数据", t)
		if t.SysDB == "" {
			fmt.Println("配置文件设置错误")
			return nil
		}
		conf = &t
	}
	return conf
}
