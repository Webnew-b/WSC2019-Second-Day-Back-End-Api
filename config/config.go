package config

import (
	"os"
	"path/filepath"
	"wscmakebygo.com/tools/fileUtil"
)

type Config struct {
	Env string `yaml:"env"`

	App   *App   `yaml:"app"`
	Log   *Log   `yaml:"log"`
	Db    *Db    `yaml:"db"`
	Redis *Redis `yaml:"redis"`
}

// 以下是2层，三层往下添加对应结构体
type App struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Log struct {
	Suffix  string `yaml:"suffix"`
	MaxSize int    `yaml:"maxSize"`
}

type Db struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
	Charset  string `yaml:"charset"`
	Loc      string `yaml:"loc"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func ReadYamlFile() []byte {
	workingPath, err := fileUtil.GetWorkingDir()
	if err != nil {
		//todo 如果文件不存在，自动生成配置文件
		panic(err)
	}
	path := filepath.Join(workingPath, "config/config.yml")
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return yamlFile
}
