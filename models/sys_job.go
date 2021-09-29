package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(SysJob))
}

type SysJob struct {
	Id      string `json:"id" orm:"pk"`
	Name    string `json:"name" valid:"Required;"`
	Enabled int8   `json:"enabled"`
	Sort    int8   `json:"sort"`
	//DeptId int64 `json:"deptId"`
	Dept *SysDept `json:"dept" orm:"column(dept_id);rel(one)"`
	BaseModel
}
