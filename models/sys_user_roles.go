package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(SysUsersRoles))
}

type SysUsersRoles struct {
	Id     string   `orm:"pk"`
	UserId *SysUser `orm:"column(user_id);rel(fk)"`
	RoleId *SysRole `orm:"column(role_id);rel(fk)"`
}
