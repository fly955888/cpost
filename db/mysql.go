package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"postman/config"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config.Mysql.Username, config.Config.Mysql.Password, config.Config.Mysql.Host, config.Config.Mysql.DB)
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接MySQL错误，原因：" + err.Error())
	} else {
		fmt.Println("连接MySQL 成功")
	}
}
