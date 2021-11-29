package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"postman/config"
	"postman/controller"
)

var (
	R *gin.Engine
)

func InitRouter() {
	gin.ForceConsoleColor()
	R = gin.New()
	R.Use(gin.Logger()).Use()
	R.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})
	R.POST("/do", controller.Do)

	admin := R.Group("/admin")

	admin.POST("/login", controller.AdminLogin)

	err := R.Run(fmt.Sprintf(":%v", config.Config.Port))
	if err != nil {
		fmt.Println("router 运行错误，原因：" + err.Error())
	}
}
