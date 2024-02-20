package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var (
	AppConf *Config
)

func Init(file string) {
	AppConf = new(Config)
	bin, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bin, AppConf)
	if err != nil {
		panic(err)
	}
}
