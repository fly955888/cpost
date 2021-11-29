package controller

import (
	"github.com/gin-gonic/gin"
	"postman/service"
)

// 初始化超级管理员信息
func Init() {
	service.InitSystemAdmin()
}

func Login(c *gin.Context) {

}
