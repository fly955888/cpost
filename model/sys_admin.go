package model

import (
	"fmt"
	"postman/db"
)

type SysAdmin struct {
	Id            int    `json:"id" gorm:"id"`
	Username      string `json:"username" gorm:"username"`
	Password      string `json:"password" gorm:"password"`
	RoleId        int    `json:"role_id" gorm:"role_id"`
	LastLoginIp   string `json:"last_login_ip" gorm:"last_login_ip"`
	LastLoginTime string `json:"last_login_time" gorm:"last_login_time"`
}

func (sa *SysAdmin) TableName() string {
	return "sys_admin"
}

func (sa *SysAdmin) GetOneByRoleId(condition int) (sysAdmin *SysAdmin) {
	db.DB.Where("role_id = ? ", condition).Find(&sysAdmin)
	fmt.Println("sysadmin=", sysAdmin)
	return
}

func (sa *SysAdmin) Insert(data SysAdmin) int64 {
	result := db.DB.Create(&data)
	return result.RowsAffected
}
