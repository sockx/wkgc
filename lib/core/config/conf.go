package config

// 其他yaml结构
/*
redis:
	address: 127.0.0.0
	password: 123a
mysql:
	username: momo
	password: 123a
*/
// 上述结构的定义方式
/*
type Config struct{
	Redis struct{
		Address string `yaml:"address"`
		Password string `yaml:"password"`
	}
	Mysql struct{
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

或者

type Config struct{
	Redis Redis `yaml:"redis"` //或者：`yaml:"redis, inline"`
	Mysql Mysql `yaml:"mysql"` //或者：`yaml:"mysql, inline"`
}
type Redis struct{
	Address string `yaml:"address"`
	Password string `yaml:"password"`
}
type Mysql struct{
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
*/

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type ConfigBase struct {
	WorkDir  string `yaml:"workdir"`
	Database string `yaml:"database"`
}

var Config *ConfigBase

func Load() error {
	conf := new(ConfigBase)
	yamlFile, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Println("tamlFile.Get err", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalln("Unmarshal:", err)
		return err
	}

	Config = conf
	return err
}
