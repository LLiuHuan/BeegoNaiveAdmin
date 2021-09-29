package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(SysMenu))
}

type SysMenu struct {
	Id            string    `json:"id" orm:"pk"`
	Name          string    `json:"name" valid:"Required;"`
	IFrame        int8      `json:"iframe"`
	Component     string    `json:"component"`
	Pid           string    `json:"pid"`
	Sort          int32     `json:"sort"`
	Icon          string    `json:"icon"`
	Path          string    `json:"path"`
	Cache         int8      `json:"cache"`
	Hidden        int8      `json:"hidden"`
	ComponentName string    `json:"componentName"`
	Permission    string    `json:"permission"`
	Type          int32     `json:"type"`
	Router        string    `json:"router"`
	RouterMethod  string    `json:"routerMethod"`
	Children      []SysMenu `json:"children" orm:"-"`
	Label         string    `orm:"-" json:"label"`
	BaseModel
}
