package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(SysRolesMenus))
}

type SysRolesMenus struct {
	Id     string   `orm:"pk"`
	MenuId *SysMenu `orm:"column(menu_id);rel(fk)"`
	RoleId *SysRole `orm:"column(role_id);rel(fk)"`
}
