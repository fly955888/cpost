package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"postman/internal"
	"postman/mdata"
)

func AdminLogin(c *gin.Context) {
	var response map[string]interface{}
	var req mdata.LoginDto
	if c.BindJSON(&req) == nil {
		if req.Username == "" || req.Password == "" {
			response["status"] = -1
			response["msg"] = "用户名或密码不能为空"
			internal.Response(c, http.StatusOK, response)
		}
	} else {
		response["status"] = -1
		response["msg"] = "json 解析失败"
		internal.Response(c, http.StatusOK, response)
	}

}
