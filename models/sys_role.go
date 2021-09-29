package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(SysRole))
}

type SysRole struct {
	Id         string     `json:"id" orm:"pk"`
	Name       string     `json:"name" valid:"Required;"`
	Remark     string     `json:"remark"`
	DataScope  string     `json:"dataScope"`
	Level      int32      `json:"level"`
	Permission string     `json:"permission"`
	Users      []*SysUser `orm:"reverse(many)"`
	Menus      []*SysMenu `json:"menus" orm:"rel(m2m);rel_through(BeegoNaiveAdmin/models.SysRolesMenus)"`
	Depts      []*SysDept `orm:"rel(m2m);rel_through(BeegoNaiveAdmin/models.SysRolesDepts)"`
	BaseModel
}
