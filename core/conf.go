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

// go get gopkg.in/yaml.v2
// 读取yaml配置文件并初始化全局配置
func InitCore() {
	// 新建一个config.Config指针, 这个是用来接收解析后的配置数据的
	c := &config.Config{}

	// 读取字节流到 yamlConf
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf err: %s", err))
	}

	// 将 yaml byte 反序列化到 指针c
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config init unmarshal: %v", err)
	}

	log.Println("config yamlFile load init success.")

	// 将解析后的配置数据赋值给全局变量
	global.Config = c
}

// 将 Config 结构体序列化为 Yaml 字节数据 并写入settings.yaml
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
