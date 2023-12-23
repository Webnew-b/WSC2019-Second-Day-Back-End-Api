package config

import (
	"os"
	"path/filepath"
	"wscmakebygo.com/tools"
)

type Config struct {
	App *App   `yaml:"app"`
	Log *Log   `yaml:"log"`
	Env string `yaml:"env"`
	Db  *Db    `yaml:"db"`
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

func ReadYamlFile() []byte {
	workingPath := tools.GetWorkingDir()
	path := filepath.Join(workingPath, "config/config.yml")
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return yamlFile
}
