package core

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"gin-blog-server/config"
	"gin-blog-server/global"
)

const ConfigFile = "settings.yaml"

// 读取yaml配置文件
// go get gopkg.in/yaml.v2
func InitCore() {

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

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}

	err = os.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("yaml配置文件修改成功")
	return nil
}
