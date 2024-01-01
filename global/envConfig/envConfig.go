package envConfig

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"wscmakebygo.com/config"
)

var (
	Config *config.Config
)

func GetConfig() config.Config {
	if Config == nil {
		panic("config not initialized")
	}
	return *Config
}

func unmarshalConfigYaml(yamlFile []byte) *config.Config {
	var _config *config.Config
	err := yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return _config
}

func InitVal() {
	log.Println("get Config")
	yamlFile := config.ReadYamlFile()
	_config := unmarshalConfigYaml(yamlFile)
	// todo 判断对应需要配置的配置值是否为空
	Config = _config
}
