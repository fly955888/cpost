package main

import (
	"fmt"
	"os"
	"postman/config"
	"postman/controller"
	"postman/db"
	"postman/router"
)

func Init() {
	args := os.Args
	if len(args) > 1 {
		conf := new(config.CONFIG)
		conf.GetConfig(os.Args[1])
	} else {
		panic("未指定配置")
	}

	db.Init()
	controller.Init()
	router.InitRouter()

}
func main() {
	fmt.Println("欢迎使用")
	Init()
	fmt.Println(config.Config)
}
