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
	"strings"
	"wkgc/lib/utils/checkerr"
	"wkgc/lib/utils/iot"

	"gopkg.in/yaml.v3"
)

type ConfigBase struct {
	WorkDir  string `yaml:"workdir"`
	Database string `yaml:"database"`
}

var Config *ConfigBase

func Load() {
	conf := new(ConfigBase)
	yamlFile, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Println("tamlFile.Get err", err)
		checkerr.CheckErr(err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalln("Unmarshal:", err)
		checkerr.CheckErr(err)
	}
	Config = conf
	// 修正结尾
	if !strings.HasSuffix(Config.WorkDir, "\\") && !strings.HasSuffix(Config.WorkDir, "/") {
		Config.WorkDir += "/"
	}
	// 拼接数据目录的绝对路径
	if strings.HasPrefix(Config.WorkDir, ".") {
		iotool := iot.NewIoTool()
		Config.WorkDir = iotool.NoWorkDir + string([]byte(Config.WorkDir)[1:len(Config.WorkDir)])
	}
	// 美化路径
	Config.WorkDir = strings.ReplaceAll(Config.WorkDir, "\\", "/")
}
