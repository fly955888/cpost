package config

import (
	"encoding/json"
	"os"
)

var (
	Config CONFIG
)

type CONFIG struct {
	Redis Redis
	Mysql Mysql
	Port  int `json:"port"`
}

type Redis struct {
	Host string `json:"host"`
	Auth string `json:"auth"`
	Db   int    `json:"db"`
}

type Mysql struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

func (conf *CONFIG) GetConfig(configPath string) {
	open, err := os.Open(configPath)
	if err != nil {
		panic("打开文件失败，原因：" + err.Error())
	}
	defer open.Close()
	decoder := json.NewDecoder(open)
	err = decoder.Decode(&Config)
	if err != nil {
		panic("解析配置文件失败，原因：" + err.Error())
	}
}
