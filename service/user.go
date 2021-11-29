package service

import (
	"fmt"
	"postman/mdata"
	"postman/model"
	"postman/utils"
	"time"
)

func InitSystemAdmin() {
	sa := new(model.SysAdmin)
	info := sa.GetOneByRoleId(0)

	if info.Id <= 0 {
		insert := sa.Insert(model.SysAdmin{Username: "admin", Password: utils.MD5("admin"), RoleId: 0, LastLoginIp: "", LastLoginTime: time.Now().Format(mdata.FormatTime)})
		fmt.Println("创建成功", insert)
	} else {
		fmt.Println("超级管理员已存在")
	}
}
