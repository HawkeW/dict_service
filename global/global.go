package global

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Sqlite Sqlite `yaml:"sqlite"`
	Http   Http   `yaml:"http"`
}

type Sqlite struct {
	Path string `yaml:"path"`
}

type Http struct {
	Port string `yaml:"port"`
}

type Opt struct {
	Config *Config
}

var Global = Opt{}

func InitConfig() {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &Global.Config)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("config.sql: %#v\n, config.http: %#v\n", Global.Config.Sqlite, Global.Config.Http)

}
