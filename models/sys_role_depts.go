package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(SysRolesDepts))
}

type SysRolesDepts struct {
	Id     string   `orm:"pk"`
	RoleId *SysRole `orm:"column(role_id);rel(fk)"`
	DeptId *SysDept `orm:"column(dept_id);rel(fk)"`
}
