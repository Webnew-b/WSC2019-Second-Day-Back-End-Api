package start

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"wscmakebygo.com/config"
	"wscmakebygo.com/global"
)

func unmarshalConfigYaml(yamlFile []byte) *config.Config {
	var _config *config.Config
	err := yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return _config
}

func createConfig() {
	yamlFile := config.ReadYamlFile()
	_config := unmarshalConfigYaml(yamlFile)
	global.Config = _config
}

func createServerAddr() string {
	serveAddr := fmt.Sprintf("%s:%d",
		global.Config.App.Host,
		global.Config.App.Port)
	return serveAddr
}
