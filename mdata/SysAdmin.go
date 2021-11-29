package mdata

type SysAdmin struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	RoleId        int    `json:"role_id"`
	LastLoginIp   string `json:"last_login_ip"`
	LastLoginTime string `json:"last_login_time"`
}

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
