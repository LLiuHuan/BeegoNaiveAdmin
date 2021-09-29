package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(SysDept))
}

type SysDept struct {
	Id       string    `json:"id" orm:"pk"`
	Name     string    `json:"name" valid:"Required;"`
	Pid      string    `json:"pid"`
	Enabled  int8      `json:"enabled" valid:"Required;"`
	Children []SysDept `orm:"-" json:"children"`
	Label    string    `orm:"-" json:"label"`
	BaseModel
}
