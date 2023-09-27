package core

import (
	"fmt"
	"gin-blog-server/config"
	"gin-blog-server/global"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// 读取yaml配置文件
// go get gopkg.in/yaml.v2
func InitCore() {
	const ConfigFile = "settings.yaml"

	c := &config.Config{}

	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf err: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config init unmarshal: %v", err)
	}

	log.Println("config yamlFile load init success.")

	global.Config = c

}
